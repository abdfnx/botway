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
      renderServiceId: ValidateProps.project.renderServiceId,
      repoBranch: ValidateProps.project.repoBranch,
      pullRequestPreviewsEnabled:
        ValidateProps.project.pullRequestPreviewsEnabled,
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
      renderServiceId,
      renderApiToken,
      icon,
      buildCommand,
      startCommand,
      rootDirectory,
      repoBranch,
      pullRequestPreviewsEnabled,
    } = req.body;

    const octokit = new Octokit({
      auth: ghToken,
    });

    const ghu = await (await octokit.request("GET /user", {})).data;

    if (!repo.toString().includes(ghu.login))
      return res.json({ message: `Repo owner must be ${ghu.login}` });

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
      renderServiceId,
      icon,
      buildCommand,
      startCommand,
      rootDirectory,
      repoBranch,
      pullRequestPreviewsEnabled,
    };

    if (hostService == "railway") {
      const getEnvId = await fetcher(
        "https://backboard.railway.app/graphql/v2",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${railwayApiToken}`,
          },
          body: JSON.stringify({
            query: `query { project(id: "${railwayProjectId}") { environments { edges { node { name, id } } } } }`,
          }),
        }
      );

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

      payload["railwayEnvId"] = envId;
    } else if (hostService == "render") {
      let vars;

      if (platform == "discord") {
        vars = [
          { key: "DISCORD_TOKEN", value: botToken },
          { key: "DISCORD_CLIENT_ID", value: botAppToken },
        ];
      } else if (platform == "slack") {
        vars = [
          { key: "SLACK_TOKEN", value: botToken },
          { key: "SLACK_APP_TOKEN", value: botAppToken },
          { key: "SLACK_SIGNING_SECRET", value: botSecretToken },
        ];
      } else if (platform == "telegram") {
        vars = [{ key: "TELEGRAM_TOKEN", value: botToken }];
      } else if (platform == "twitch") {
        vars = [
          { key: "TWITCH_OAUTH_TOKEN", value: botToken },
          { key: "TWITCH_CLIENT_ID", value: botAppToken },
          { key: "TWITCH_CLIENT_SECRET", value: botSecretToken },
        ];
      }

      await fetcher(
        `https://api.render.com/v1/services/${renderServiceId}/env-vars`,
        {
          method: "PUT",
          headers: {
            Accept: "application/json",
            "Content-Type": "application/json",
            Authorization: `Bearer ${renderApiToken}`,
          },
          body: JSON.stringify(vars),
        }
      );
    }

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
