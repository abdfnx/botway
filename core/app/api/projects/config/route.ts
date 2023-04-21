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

  const { payload: bt } = await jwtDecrypt(body.botToken, BW_SECRET_KEY);

  let bat: any, bst: any;

  if (body.botAppToken) {
    const { payload } = await jwtDecrypt(body.botAppToken, BW_SECRET_KEY);

    bat = payload;
  } else {
    bat = "not";
  }

  if (body.botSecretToken) {
    const { payload } = await jwtDecrypt(body.botSecretToken, BW_SECRET_KEY);

    bst = payload;
  } else {
    bst = "not";
  }

  const { payload: railwayProjectId } = await jwtDecrypt(
    body.railwayProjectId,
    BW_SECRET_KEY
  );

  const { payload: railwayServiceId } = await jwtDecrypt(
    body.railwayServiceId,
    BW_SECRET_KEY
  );

  const getEnvId = await fetcher("https://backboard.railway.app/graphql/v2", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${railwayApiToken.data}`,
    },
    body: JSON.stringify({
      query: `
        query {
          project(id: "${railwayProjectId.data}") {
            environments {
              edges {
                node {
                  name,
                  id
                }
              }
            }
          }
        }
      `,
    }),
  });

  if (getEnvId.errors) {
    return NextResponse.json({ error: getEnvId.errors[0].message });
  }

  const envId = getEnvId.data.project.environments.edges.find(
    (env: any) => env.node.name == "production"
  ).node.id;

  let vars;

  if (body.platform == "discord") {
    vars = `DISCORD_TOKEN: "${bt.data}", DISCORD_CLIENT_ID: "${bat.data}"`;
  } else if (body.platform == "slack") {
    vars = `SLACK_TOKEN: "${bt.data}", SLACK_APP_TOKEN: "${bat.data}", SLACK_SIGNING_SECRET: "${bst.data}"`;
  } else if (body.platform == "telegram") {
    vars = `TELEGRAM_TOKEN: "${bt.data}"`;
  } else if (body.platform == "twitch") {
    vars = `TWITCH_OAUTH_TOKEN: "${bt.data}", TWITCH_CLIENT_ID: "${bat.data}", TWITCH_CLIENT_SECRET: "${bst.data}"`;
  }

  const octokit = new Octokit({
    auth: githubApiToken.data,
  });

  const ghu = await (await octokit.request("GET /user", {})).data;

  if (!body.repo.toString().includes(ghu.login))
    NextResponse.json({ error: `Repo owner must be ${ghu.login}` });

  await fetcher("https://backboard.railway.app/graphql/v2", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${railwayApiToken.data}`,
    },
    body: JSON.stringify({
      operationName: "setTokens",
      query: `
        mutation setTokens {
          variableCollectionUpsert(input: {
            projectId: "${railwayProjectId.data}",
            environmentId: "${envId}",
            serviceId: "${railwayServiceId.data}",
            variables: { ${vars} }
          })
        }
      `,
    }),
  });

  let updateBody = {
    bot_token: body.botToken,
    bot_app_token: "",
    bot_secret_token: "",
  };

  if (body.platform != "telegram") {
    updateBody["bot_app_token"] = body.botAppToken;
  }

  if (body.platform == "slack" || body.platform == "twitch") {
    updateBody["bot_secret_token"] = body.botSecretToken;
  }

  const { error } = await supabase
    .from("projects")
    .update(updateBody)
    .eq("id", body.projectId);

  if (error) {
    return NextResponse.json({ error });
  }

  return NextResponse.json({ message: "Success" });
}
