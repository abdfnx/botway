import { NextResponse } from "next/server";
import { jwtDecrypt } from "jose";
import { BW_SECRET_KEY } from "@/tools/tokens";
import { fetcher } from "@/tools/fetch";
import createClient from "@/supabase/server";

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

  const { payload: zeaburApiToken } = await jwtDecrypt(
    user?.user_metadata["zeaburApiToken"],
    BW_SECRET_KEY,
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

  const { payload: zeaburServiceId } = await jwtDecrypt(
    body.zeaburServiceId,
    BW_SECRET_KEY,
  );

  const { payload: zeaburEnvId } = await jwtDecrypt(
    body.zeaburEnvId,
    BW_SECRET_KEY,
  );

  let vars;
  let shared = `createEnvironmentVariable(
        environmentID: "${zeaburEnvId.data}"
        serviceID: "${zeaburServiceId.data}"`;

  if (body.platform === "discord") {
    vars = `
      ${shared}
        key: "DISCORD_TOKEN"
        value: "${bt.data}"
      ) {
        _id
      }

      var2: ${shared}
        key: "DISCORD_CLIENT_ID"
        value: "${bat.data}"
      ) {
        _id
      }
    `;
  } else if (body.platform === "slack") {
    vars = `
      ${shared}
        key: "SLACK_TOKEN"
        value: "${bt.data}"
      ) {
        _id
      }

      var2: ${shared}
        key: "SLACK_APP_TOKEN"
        value: "${bat.data}"
      ) {
        _id
      }

      var3: ${shared}
        key: "SLACK_SIGNING_SECRET"
        value: "${bst.data}"
      ) {
        _id
      }
    `;
  } else if (body.platform === "telegram") {
    vars = `
      ${shared}
        key: "TELEGRAM_TOKEN"
        value: "${bt.data}"
      ) {
        _id
      }
    `;
  } else if (body.platform === "twitch") {
    vars = `
      ${shared}
        key: "TWITCH_OAUTH_TOKEN"
        value: "${bt.data}"
      ) {
        _id
      }

      var2: ${shared}
        key: "TWITCH_CLIENT_ID"
        value: "${bat.data}"
      ) {
        _id
      }

      var3: ${shared}
        key: "TWITCH_CLIENT_SECRET"
        value: "${bst.data}"
      ) {
        _id
      }
    `;
  }

  console.log(vars);

  await fetcher("https://gateway.zeabur.com/graphql", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${zeaburApiToken.data}`,
    },
    body: JSON.stringify({
      query: `
        mutation {
          ${vars}
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

  if (body.platform === "slack" || body.platform === "twitch") {
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
