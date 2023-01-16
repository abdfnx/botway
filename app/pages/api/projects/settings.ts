import { ncOpts } from "@/api/nc";
import nc from "next-connect";
import { fetcher } from "@/lib/fetch";
import { auths } from "@/api/middlewares";
import multer from "multer";
import { getMongoDb } from "@/api/mongodb";
import { updateProject } from "@/api/db";
import { jwtDecrypt } from "jose";
import { BW_SECRET_KEY } from "@/tools/api-tokens";

const handler = nc(ncOpts);

handler.use(...auths);

handler.patch(multer({ dest: "/tmp" }).single("data"), async (req, res) => {
  if (!req.user) {
    return res.status(401).end();
  }

  if (!req.user.emailVerified && process.env.NEXT_PUBLIC_FULL == "true") {
    return res.status(401).json({ message: "You must verify your email" });
  }

  const db = await getMongoDb();

  let {
    id,
    name,
    repo,
    userId,
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
    railwayEnvId,
    icon,
    buildCommand,
    startCommand,
    rootDirectory,
    repoBranch,
    pullRequestPreviewsEnabled,
  } = req.body;

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

    const nameBody =
      name != ""
        ? `projectUpdate(id: "${rwProjectId.data}", input: { name: "${name}" }) { id }`
        : "";
    const serviceNameBody = name != "" ? `name: "${name}-main"` : "";
    const iconBody = icon != "" ? `icon: "${icon}"` : "";
    const repoBody = repo != "" ? `source: { repo: "${repo}" }` : "";
    const buildCommandBody =
      buildCommand != "" ? `buildCommand: "${buildCommand}"` : "";
    const rootDirectoryBody =
      rootDirectory != "" ? `rootDirectory: "${rootDirectory}"` : "";
    const startCommandBody =
      startCommand != "" ? `startCommand: "${startCommand}"` : "";

    const query = `mutation settingsUpdate { ${nameBody} serviceUpdate(id: "${rwServiceId.data}", input: { ${serviceNameBody} ${iconBody} ${buildCommandBody} ${rootDirectoryBody} ${startCommandBody} ${repoBody} }) { id }}`;

    await fetcher("https://backboard.railway.app/graphql/v2", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${rwApiToken.data}`,
      },
      body: JSON.stringify({
        operationName: "settingsUpdate",
        query,
      }),
    });
  } else if (hostService == "render") {
    const { payload: rndApiToken } = await jwtDecrypt(
      renderApiToken,
      BW_SECRET_KEY
    );
    const { payload: rndServiceId } = await jwtDecrypt(
      renderServiceId,
      BW_SECRET_KEY
    );

    const body = JSON.stringify({
      serviceDetails: { pullRequestPreviewsEnabled },
      branch: repoBranch,
      name,
    });

    await fetcher(`https://api.render.com/v1/services/${rndServiceId.data}`, {
      method: "PATCH",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
        Authorization: `Bearer ${rndApiToken.data}`,
      },
      body,
    });
  }

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
    railwayEnvId,
    renderServiceId,
    icon,
    buildCommand,
    startCommand,
    rootDirectory,
    repoBranch,
    pullRequestPreviewsEnabled,
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
});

export const config = {
  api: {
    bodyParser: false,
  },
};

export default handler;
