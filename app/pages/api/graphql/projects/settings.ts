import { ncOpts } from "@/api/nc";
import nc from "next-connect";
import { fetcher } from "@/lib/fetch";
import { auths } from "@/api/middlewares";
import multer from "multer";
import { getMongoDb } from "@/api/mongodb";
import { updateProject } from "@/api/db";

const handler = nc(ncOpts);

handler.use(...auths);

handler.patch(multer({ dest: "/tmp" }).single("data"), async (req, res) => {
  if (!req.user) {
    return res.status(401).end();
  }

  const db = await getMongoDb();

  const {
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
    renderProjectId,
    railwayEnvId,
    icon,
    buildCommand,
    startCommand,
    rootDirectory,
  } = req.body;

  const nameBody =
    name != ""
      ? `projectUpdate(id: "${railwayProjectId}", input: { name: "${name}" }) { id }`
      : "";
  const serviceNameBody = name != "" ? `name: "${name}-main"` : "";
  const iconBody = icon != "" ? `icon: "${icon}"` : "";
  const repoBody = repo != "" ? `source { repo: "${repo}" }` : "";
  const buildCommandBody =
    buildCommand != "" ? `buildCommand: "${buildCommand}"` : "";
  const rootDirectoryBody =
    rootDirectory != "" ? `rootDirectory: "${rootDirectory}"` : "";
  const startCommandBody =
    startCommand != "" ? `startCommand: "${startCommand}"` : "";

  const query = `mutation settingsUpdate { ${nameBody} serviceUpdate(id: "${railwayServiceId}", input: { ${serviceNameBody} ${iconBody} ${buildCommandBody} ${rootDirectoryBody} ${startCommandBody} ${repoBody} }) { id }}`;

  await fetcher("https://backboard.railway.app/graphql/v2", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${railwayApiToken}`,
    },
    body: JSON.stringify({
      operationName: "settingsUpdate",
      query,
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
    railwayEnvId,
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
});

export const config = {
  api: {
    bodyParser: false,
  },
};

export default handler;
