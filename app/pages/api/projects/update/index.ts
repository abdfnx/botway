import { ValidateProps } from "@/api/constants";
import { updateProject } from "@/api/db";
import { auths, validateBody } from "@/api/middlewares";
import { getMongoDb } from "@/api/mongodb";
import { ncOpts } from "@/api/nc";
import { fetcher } from "@/lib/fetch";
import multer from "multer";
import nc from "next-connect";
import { Octokit } from "octokit";

const handler = nc(ncOpts);

handler.use(...auths);

handler.patch(
  multer({ dest: "/tmp" }).single("data"),
  validateBody({
    type: "object",
    properties: {
      name: ValidateProps.project.name,
      platform: ValidateProps.project.platform,
      botToken: ValidateProps.project.botToken,
      botAppToken: ValidateProps.project.botAppToken,
      botSecretToken: ValidateProps.project.botSecretToken,
      railwayProjectId: ValidateProps.project.railwayProjectId,
      renderProjectId: ValidateProps.project.renderProjectId,
    },
    additionalProperties: true,
  }),
  async (req, res) => {
    if (!req.user) {
      req.status(401).end();

      return;
    }

    const db = await getMongoDb();

    let {
      id,
      name,
      ghToken,
      botToken,
      platform,
      botAppToken,
      botSecretToken,
      railwayApiToken,
      railwayProjectId,
      railwayServiceId,
    } = req.body;

    const getEnvId = await fetcher("https://backboard.railway.app/graphql/v2", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${railwayApiToken}`,
      },
      body: JSON.stringify({
        query: `query { project(id: "${railwayProjectId}") { environments { edges { node { name, id } } } } }`,
      }),
    });

    const envId = getEnvId.data.project.environments.edges.find(
      (env: any) => env.node.name == "production"
    ).node.id;

    let vars;

    if (platform == "discord") {
      vars = `DISCORD_TOKEN: "${botToken}", DISCORD_CLIENT_ID: "${botAppToken}"`;
    } else if (platform == "slack") {
      vars = `SLACK_TOKEN: "${botToken}", SLACK_APP_TOKEN: "${botAppToken}", SLACK_SIGNING_SECRET: "${botSecretToken}"`;
    } else if (platform == "telegram") {
      vars = `TELEGRAM_TOKEN: "${botToken}"`;
    } else if (platform == "twitch") {
      vars = `TWITCH_OAUTH_TOKEN: "${botToken}", TWITCH_CLIENT_ID: "${botAppToken}", TWITCH_CLIENT_SECRET: "${botSecretToken}"`;
    }

    const octokit = new Octokit({
      auth: ghToken,
    });

    const ghu = await (await octokit.request("GET /user", {})).data;

    await fetcher("https://backboard.railway.app/graphql/v2", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${railwayApiToken}`,
      },
      body: JSON.stringify({
        operationName: "setTokens",
        query: `mutation setTokens { variableCollectionUpsert(input: { projectId: "${railwayProjectId}", environmentId: "${envId}", serviceId: "${railwayServiceId}", variables: { ${vars} } }) serviceUpdate(id: "${railwayServiceId}", input: { source: { repo: "${ghu.login}/${name}" } }) { source { repo } } }`,
      }),
    });

    let payload = {
      ...(name && { name }),
      botToken,
      railwayEnvId: envId,
    };

    if (platform != "telegram") {
      payload = {
        ...(name && { name }),
        botToken,
        botAppToken,
        railwayEnvId: envId,
      };
    } else if (platform == "slack" || platform == "twitch") {
      payload = {
        ...(name && { name }),
        botToken,
        botAppToken,
        botSecretToken,
        railwayEnvId: envId,
      };
    }

    const prj = await updateProject(db, id, payload);

    res.json({ prj });
  }
);

export const config = {
  api: {
    bodyParser: false,
  },
};

export default handler;
