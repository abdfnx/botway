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
      icon: ValidateProps.project.icon,
      buildCommand: ValidateProps.project.buildCommand,
      startCommand: ValidateProps.project.startCommand,
      rootDirectory: ValidateProps.project.rootDirectory,
      visibility: ValidateProps.project.visibility,
      platform: ValidateProps.project.platform,
      lang: ValidateProps.project.lang,
      packageManager: ValidateProps.project.packageManager,
      hostService: ValidateProps.project.hostService,
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
      repo,
      userId,
      ghToken,
      visibility,
      lang,
      packageManager,
      hostService,
      botToken,
      platform,
      botAppToken,
      botSecretToken,
      railwayApiToken,
      railwayProjectId,
      railwayServiceId,
      renderProjectId,
      icon,
      buildCommand,
      startCommand,
      rootDirectory,
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

    if (!repo.includes(ghu.login))
      return res.json({ message: `Repo owner must be ${ghu.login}` });

    const repoBody =
      repo != ""
        ? `source: { repo: "${repo}" }`
        : `source: { repo: "${ghu.login}/${name}" }`;

    await fetcher("https://backboard.railway.app/graphql/v2", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${railwayApiToken}`,
      },
      body: JSON.stringify({
        operationName: "setTokens",
        query: `mutation setTokens { variableCollectionUpsert(input: { projectId: "${railwayProjectId}", environmentId: "${envId}", serviceId: "${railwayServiceId}", variables: { ${vars} } }) serviceUpdate(id: "${railwayServiceId}", input: { ${repoBody} }) { source { repo } } }`,
      }),
    });

    let payload = {
      id,
      ...(name && { name }),
      repo,
      botToken,
      platform,
      lang,
      packageManager,
      visibility,
      hostService,
      railwayProjectId,
      railwayServiceId,
      railwayEnvId: envId,
      renderProjectId,
      icon,
      buildCommand,
      startCommand,
      rootDirectory,
    };

    if (platform != "telegram") {
      payload["botToken"] = botToken;
      payload["botAppToken"] = botAppToken;
    } else if (platform == "slack" || platform == "twitch") {
      payload["botToken"] = botToken;
      payload["botAppToken"] = botAppToken;
      payload["botSecretToken"] = botSecretToken;
    }

    const prj = await updateProject(db, userId, id, payload);

    res.json({ prj });
  }
);

export const config = {
  api: {
    bodyParser: false,
  },
};

export default handler;
