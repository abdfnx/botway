import { ValidateProps } from "@/api/constants";
import { updateProject } from "@/api/db";
import { auths, validateBody } from "@/api/middlewares";
import { getMongoDb } from "@/api/mongodb";
import { ncOpts } from "@/api/nc";
import { fetcher } from "@/lib/fetch";
import { BW_SECRET_KEY } from "@/tools/api-tokens";
import multer from "multer";
import nc from "next-connect";
import { Octokit } from "octokit";
import { EncryptJWT, jwtDecrypt } from "jose";

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
      rwEnvId,
      renderServiceId,
      renderApiToken,
      icon,
      buildCommand,
      startCommand,
      rootDirectory,
      repoBranch,
      pullRequestPreviewsEnabled,
    } = req.body;

    const { payload: ghApiToken } = await jwtDecrypt(ghToken, BW_SECRET_KEY);

    const octokit = new Octokit({
      auth: ghApiToken.data,
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

    const { payload: bt } = await jwtDecrypt(botToken, BW_SECRET_KEY);

    let bat: any, bst: any;

    if (botAppToken) {
      const { payload } = await jwtDecrypt(botAppToken, BW_SECRET_KEY);

      bat = payload;
    } else {
      bat = "not";
    }

    if (botSecretToken) {
      const { payload } = await jwtDecrypt(botSecretToken, BW_SECRET_KEY);

      bst = payload;
    } else {
      bst = "not";
    }

    if (hostService == "railway") {
      const { payload: rwApiToken } = await jwtDecrypt(
        railwayApiToken,
        BW_SECRET_KEY
      );
      const { payload: rwProjectId } = await jwtDecrypt(
        railwayProjectId,
        BW_SECRET_KEY
      );
      const { payload: rwServiceId } = await jwtDecrypt(
        railwayServiceId,
        BW_SECRET_KEY
      );

      const getEnvId = await fetcher(
        "https://backboard.railway.app/graphql/v2",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${rwApiToken.data}`,
          },
          body: JSON.stringify({
            query: `query { project(id: "${rwProjectId.data}") { environments { edges { node { name, id } } } } }`,
          }),
        }
      );

      const envId = getEnvId.data.project.environments.edges.find(
        (env: any) => env.node.name == "production"
      ).node.id;

      let vars;

      if (platform == "discord") {
        vars = `DISCORD_TOKEN: "${bt.data}", DISCORD_CLIENT_ID: "${bat.data}"`;
      } else if (platform == "slack") {
        vars = `SLACK_TOKEN: "${bt.data}", SLACK_APP_TOKEN: "${bat.data}", SLACK_SIGNING_SECRET: "${bst.data}"`;
      } else if (platform == "telegram") {
        vars = `TELEGRAM_TOKEN: "${bt.data}"`;
      } else if (platform == "twitch") {
        vars = `TWITCH_OAUTH_TOKEN: "${bt.data}", TWITCH_CLIENT_ID: "${bat.data}", TWITCH_CLIENT_SECRET: "${bst.data}"`;
      }

      const repoBody =
        repo != ""
          ? `source: { repo: "${repo}" }`
          : `source: { repo: "${ghu.login}/${name}" }`;

      await fetcher("https://backboard.railway.app/graphql/v2", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${rwApiToken.data}`,
        },
        body: JSON.stringify({
          operationName: "setTokens",
          query: `mutation setTokens { variableCollectionUpsert(input: { projectId: "${rwProjectId.data}", environmentId: "${envId}", serviceId: "${rwServiceId.data}", variables: { ${vars} } }) serviceUpdate(id: "${rwServiceId.data}", input: { ${repoBody} }) { source { repo } } }`,
        }),
      });

      const railwayEnvId = await new EncryptJWT({
        data: envId,
      })
        .setProtectedHeader({ alg: "dir", enc: "A128CBC-HS256" })
        .encrypt(BW_SECRET_KEY);

      payload["railwayEnvId"] = railwayEnvId;
    } else if (hostService == "render") {
      const { payload: rndApiToken } = await jwtDecrypt(
        renderApiToken,
        BW_SECRET_KEY
      );
      const { payload: rndServiceId } = await jwtDecrypt(
        renderServiceId,
        BW_SECRET_KEY
      );

      let vars;

      if (platform == "discord") {
        vars = [
          { key: "DISCORD_TOKEN", value: bt.data },
          { key: "DISCORD_CLIENT_ID", value: bat.data },
        ];
      } else if (platform == "slack") {
        vars = [
          { key: "SLACK_TOKEN", value: bt.data },
          { key: "SLACK_APP_TOKEN", value: bat.data },
          { key: "SLACK_SIGNING_SECRET", value: bst.data },
        ];
      } else if (platform == "telegram") {
        vars = [{ key: "TELEGRAM_TOKEN", value: bt.data }];
      } else if (platform == "twitch") {
        vars = [
          { key: "TWITCH_OAUTH_TOKEN", value: bt.data },
          { key: "TWITCH_CLIENT_ID", value: bat.data },
          { key: "TWITCH_CLIENT_SECRET", value: bst.data },
        ];
      }

      await fetcher(
        `https://api.render.com/v1/services/${rndServiceId.data}/env-vars`,
        {
          method: "PUT",
          headers: {
            Accept: "application/json",
            "Content-Type": "application/json",
            Authorization: `Bearer ${rndApiToken.data}`,
          },
          body: JSON.stringify(vars),
        }
      );
    }

    if (platform != "telegram") {
      payload["botToken"] = bt.data;
      payload["botAppToken"] = bat.data;
    } else if (platform == "slack" || platform == "twitch") {
      payload["botToken"] = bt.data;
      payload["botAppToken"] = bat.data;
      payload["botSecretToken"] = bst.data;
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
