import { NextResponse } from "next/server";
import { fetcher } from "@/tools/fetch";
import { jwtDecrypt } from "jose";
import { BW_SECRET_KEY } from "@/tools/tokens";
import createClient from "@/supabase/server";
import { Octokit } from "octokit";
import { faker } from "@faker-js/faker";

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

  const { payload: railwayApiToken } = await jwtDecrypt(
    user?.user_metadata["railwayApiToken"],
    BW_SECRET_KEY,
  );

  const { payload: password } = await jwtDecrypt(body.password, BW_SECRET_KEY);

  const { data, error: ceError } = await supabase
    .from("projects")
    .select("railway_project_id, lang")
    .eq("id", body.projectId)
    .single();

  if (ceError) {
    return NextResponse.json({ error: ceError });
  }

  const { payload: railwayProjectId } = await jwtDecrypt(
    data.railway_project_id,
    BW_SECRET_KEY,
  );

  const octokit = new Octokit({
    auth: githubApiToken.data,
  });

  const ghu = await (await octokit.request("GET /user", {})).data;

  let pkgs = "cmake ";

  switch (data.lang) {
    case "crystal":
      pkgs += "crystal";

      break;

    case "csharp":
      pkgs += "dotnet";

      break;

    case "dart":
      pkgs += "dart-lang/dart/dart";

      break;

    case "deno":
      pkgs += "deno";

      break;

    case "go":
      pkgs += "go";

      break;

    case "java":
      pkgs += "java gradle";

      break;
    case "kotlin":
      pkgs += "kotlin java gradle";

      break;

    case "nim":
      pkgs += "nim";

      break;

    case "php":
      pkgs += "composer";

      break;

    case "python":
      pkgs += "poetry pipenv";

      break;

    case "rust":
      pkgs += "rust rustup-init";

      break;

    case "swift":
      pkgs += "swift";

      break;
  }

  const query = `
    mutation {
      templateDeploy(input: {
        services: [
          {
            hasDomain: true
            isPrivate: true
            owner: "${ghu.login}"
            name: "bwce-${body.slug}-${faker.number.int({ max: 100 })}"
            serviceName: "CE"
            template: "https://github.com/botwayorg/ce"
            variables: {
              GIT_REPO: "https://github.com/${body.repo}"
              GITHUB_TOKEN: "${githubApiToken.data}"
              PASSWORD: "${password.data}"
              PKGS: "${pkgs}"
            }
            volumes: [
              {
                mountPath: "/root"
                projectId: "${railwayProjectId.data}"
              }
            ]
          }
        ]

        projectId: "${railwayProjectId.data}"
      }) {
        projectId
      }
    }
  `;

  const enable = await fetcher("https://backboard.railway.app/graphql/v2", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${railwayApiToken.data}`,
    },
    body: JSON.stringify({
      query,
    }),
  });

  if (enable.errors) {
    console.log(enable.errors);

    return NextResponse.json({ message: enable.errors[0].message });
  }

  const { error } = await supabase
    .from("projects")
    .update({
      enable_ce: true,
    })
    .eq("id", body.projectId);

  if (error) {
    return NextResponse.json({ error });
  }

  return NextResponse.json({ message: "Success" });
}
