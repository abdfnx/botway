import { NextResponse } from "next/server";
import { fetcher } from "@/tools/fetch";
import { EncryptJWT, jwtDecrypt } from "jose";
import { BW_SECRET_KEY } from "@/tools/tokens";
import createClient from "@/supabase/server";
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

  const { payload: zeaburApiToken } = await jwtDecrypt(
    user?.user_metadata["zeaburApiToken"],
    BW_SECRET_KEY,
  );

  const { payload: password } = await jwtDecrypt(body.password, BW_SECRET_KEY);

  const { data, error: ceError } = await supabase
    .from("projects")
    .select("zeabur_project_id, zeabur_env_id, lang, repo")
    .eq("id", body.projectId)
    .single();

  if (ceError) {
    return NextResponse.json({ error: ceError });
  }

  const { payload: zeaburProjectId } = await jwtDecrypt(
    data.zeabur_project_id,
    BW_SECRET_KEY,
  );

  const { payload: zeaburEnvId } = await jwtDecrypt(
    data.zeabur_env_id,
    BW_SECRET_KEY,
  );

  let pkgs = "go ";

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

    case "kotlin":
      pkgs += "kotlin";

      break;

    case "nim":
      pkgs += "nim";

      break;

    case "php":
      pkgs += "composer";

      break;

    case "python":
      pkgs += "poetry";

      break;

    case "swift":
      pkgs += "swift";

      break;
  }

  const query = `
    mutation {
      createServiceFromMarketplace(
        projectID: "${zeaburProjectId.data}"
        itemCode: "botway-ce"
        name: "BotwayCE"
      ) {
        _id
        createdAt
      }
    }
  `;

  const enable = await fetcher("https://gateway.zeabur.com/graphql", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${zeaburApiToken.data}`,
    },
    body: JSON.stringify({
      query,
    }),
  });

  if (enable.errors) {
    return NextResponse.json({ message: enable.errors[0].message });
  } else {
    const query = `
      mutation {
        updateEnvironmentVariable(
          serviceID: "${enable.data.createServiceFromMarketplace._id}"
          environmentID: "${zeaburEnvId.data}"
          data: {
            GIT_REPO: "https://github.com/${data.repo}"
            GH_TOKEN: "${githubApiToken.data}"
            PASSWORD: "${password.data}"
            PKGS: "${pkgs}"
          }
        )

        addDomain(
          serviceID: "${enable.data.createServiceFromMarketplace._id}"
          environmentID: "${zeaburEnvId.data}"
          domain: "${body.slug + "-" + faker.word.sample()}"
          isGenerated: true
        ) {
          _id
          domain
        }
      }
    `;

    const update = await fetcher("https://gateway.zeabur.com/graphql", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${zeaburApiToken.data}`,
      },
      body: JSON.stringify({
        query,
      }),
    });

    if (update.errors) {
      return NextResponse.json({ message: update.errors[0].message });
    }
  }

  const ce_service_id = await new EncryptJWT({
    data: enable.data.createServiceFromMarketplace._id,
  })
    .setProtectedHeader({ alg: "dir", enc: "A128CBC-HS256" })
    .encrypt(BW_SECRET_KEY);

  const { error } = await supabase
    .from("projects")
    .update({
      enable_ce: true,
      ce_service_id,
    })
    .eq("id", body.projectId);

  if (error) {
    return NextResponse.json({ error });
  }

  return NextResponse.json({ message: "Success" });
}
