import { NextResponse } from "next/server";
import { EncryptJWT, jwtDecrypt } from "jose";
import { BW_SECRET_KEY } from "@/tools/tokens";
import { fetcher } from "@/tools/fetch";
import { Octokit } from "octokit";
import createClient from "@/supabase/server";
import { exec } from "child_process";
import { stringify } from "ajv";

export const revalidate = 0;

export async function POST(request: Request) {
  const body = await request.json();

  const supabase = createClient();

  const {
    data: { user: user },
    error: userError,
  } = await supabase.auth.getUser();

  if (userError) {
    return NextResponse.json({ error: userError });
  }

  const { payload: githubApiToken } = await jwtDecrypt(
    user?.user_metadata["githubApiToken"],
    BW_SECRET_KEY,
  );

  const { payload: zeaburApiToken } = await jwtDecrypt(
    user?.user_metadata["zeaburApiToken"],
    BW_SECRET_KEY,
  );

  const octokit = new Octokit({
    auth: githubApiToken.data,
  });

  const ghu = await (await octokit.request("GET /user", {})).data;

  const repo = await octokit.request("POST /user/repos", {
    name: body.name,
    description: `My Awesome Bot. Created By @botwayorg, Hosted On @zeabur.`,
    private: body.visibility != "public",
  });

  exec(
    `create-botway-bot ${stringify(body.name)} ${stringify(
      body.platform,
    )} ${stringify(body.lang)} ${stringify(body.package_manager)} ${stringify(
      githubApiToken.data,
    )} ${stringify(ghu.login)} ${stringify(ghu.email)}`,
  )
    .on("error", (error) => {
      console.log(error.message);

      return NextResponse.json({ error: error.message });
    })
    .on("message", (m) => {
      console.log(m.toString());
    });

  const createProject = await fetcher("https://gateway.zeabur.com/graphql", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${zeaburApiToken.data}`,
    },
    body: JSON.stringify({
      query: `
        mutation {
          createProject(name: "${body.name}") {
            _id
            environments {
              _id
            }
          }
        }
      `,
    }),
  });

  if (createProject.errors) {
    console.log(createProject.errors);
    console.log(createProject.errors[0].message);

    return NextResponse.json({ error: createProject.errors[0].message });
  }

  const createService = await fetcher("https://gateway.zeabur.com/graphql", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${zeaburApiToken.data}`,
    },
    body: JSON.stringify({
      query: `
        mutation {
          createService(
            name: "main"
            template: GIT
            projectID: "${createProject.data.createProject._id}"
            gitProvider: GITHUB
            repoID: ${repo.data.id}
            branchName: "main"
          ) {
            _id
          }
        }
      `,
    }),
  });

  if (createService.errors) {
    console.log(createService.errors);
    console.log(createService.errors[0].message);

    return NextResponse.json({ error: createService.errors[0].message });
  }

  const zeaburProjectId = await new EncryptJWT({
    data: createProject.data.createProject._id,
  })
    .setProtectedHeader({ alg: "dir", enc: "A128CBC-HS256" })
    .encrypt(BW_SECRET_KEY);

  const zeaburEnvId = await new EncryptJWT({
    data: createProject.data.createProject.environments[0]._id,
  })
    .setProtectedHeader({ alg: "dir", enc: "A128CBC-HS256" })
    .encrypt(BW_SECRET_KEY);

  const zeaburServiceId = await new EncryptJWT({
    data: createService.data.createService._id,
  })
    .setProtectedHeader({ alg: "dir", enc: "A128CBC-HS256" })
    .encrypt(BW_SECRET_KEY);

  const { error } = await supabase.from("projects").insert({
    user_id: user?.id,
    name: body.name,
    repo: `${ghu.login}/${body.name}`,
    platform: body.platform,
    lang: body.lang,
    package_manager: body.package_manager,
    icon: "",
    root_directory: "",
    bot_token: "",
    bot_app_token: "",
    bot_secret_token: "",
    zeabur_project_id: zeaburProjectId,
    zeabur_env_id: zeaburEnvId,
    zeabur_service_id: zeaburServiceId,
    build_command: "",
    start_command: "",
  });

  if (error) {
    console.log(error.message);

    return NextResponse.json({ error });
  }

  return NextResponse.json({ message: "Success" });
}
