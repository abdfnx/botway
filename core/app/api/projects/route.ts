import { NextResponse } from "next/server";
import { EncryptJWT, jwtDecrypt } from "jose";
import { BW_SECRET_KEY } from "@/tools/tokens";
import { fetcher } from "@/tools/fetch";
import { Octokit } from "octokit";
import createClient from "@/supabase/server";
import { exec } from "child_process";
import { stringify } from "ajv";

// export const runtime = "edge";
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
    BW_SECRET_KEY
  );

  const { payload: railwayApiToken } = await jwtDecrypt(
    user?.user_metadata["railwayApiToken"],
    BW_SECRET_KEY
  );

  const createRailwayProject = await fetcher(
    "https://backboard.railway.app/graphql/v2",
    {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${railwayApiToken.data}`,
      },
      body: JSON.stringify({
        operationName: "projectCreate",
        query: `
          mutation projectCreate {
            projectCreate(input: {
              name: "${body.name}"
            }) {
              id
            }
          }
        `,
      }),
    }
  );

  if (createRailwayProject.errors) {
    return NextResponse.json({ error: createRailwayProject.errors[0].message });
  }

  const createService = await fetcher(
    "https://backboard.railway.app/graphql/v2",
    {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${railwayApiToken.data}`,
      },
      body: JSON.stringify({
        operationName: "serviceCreate",
        query: `
          mutation serviceCreate {
            serviceCreate(input: {
              name: "main",
              projectId: "${createRailwayProject.data.projectCreate.id}"
            }) {
              id
            }
          }
        `,
      }),
    }
  );

  if (createService.errors) {
    return NextResponse.json({ error: createService.errors[0].message });
  }

  const railwayProjectId = await new EncryptJWT({
    data: createRailwayProject.data.projectCreate.id,
  })
    .setProtectedHeader({ alg: "dir", enc: "A128CBC-HS256" })
    .encrypt(BW_SECRET_KEY);

  const railwayServiceId = await new EncryptJWT({
    data: createService.data.serviceCreate.id,
  })
    .setProtectedHeader({ alg: "dir", enc: "A128CBC-HS256" })
    .encrypt(BW_SECRET_KEY);

  const octokit = new Octokit({
    auth: githubApiToken.data,
  });

  const ghu = await (await octokit.request("GET /user", {})).data;

  await octokit.request("POST /user/repos", {
    name: body.name,
    description: `My Awesome ${body.platform} botway bot.`,
    private: body.visibility != "public",
  });

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
    railway_project_id: railwayProjectId,
    railway_service_id: railwayServiceId,
    build_command: "",
    start_command: "",
  });

  if (error) {
    return NextResponse.json({ error });
  }

  exec(
    `create-botway-bot ${stringify(body.name)} ${stringify(
      body.platform
    )} ${stringify(body.lang)} ${stringify(
      body.package_manager
    )} railway ${stringify(githubApiToken.data)} ${stringify(
      ghu.login
    )} ${stringify(ghu.email)}`
  )
    .on("error", (error) => {
      return NextResponse.json({ error: error.message });
    })
    .on("message", (m) => {
      console.log(m.toString());
    });

  return NextResponse.json({ message: "Success" });
}
