import { ValidateProps } from "@/api/constants";
import { insertProject } from "@/api/db";
import { auths, validateBody } from "@/api/middlewares";
import { getMongoDb } from "@/api/mongodb";
import { ncOpts } from "@/api/nc";
import { fetcher } from "@/lib/fetch";
import { BW_SECRET_KEY } from "@/tools/api-tokens";
import nc from "next-connect";
import { Octokit } from "octokit";
import { EncryptJWT, jwtDecrypt } from "jose";
import { exec } from "child_process";
import { stringify } from "ajv";

const handler = nc(ncOpts);

handler.post(
  ...auths,
  validateBody({
    type: "object",
    properties: {
      name: ValidateProps.project.name,
      icon: ValidateProps.project.icon,
      repo: ValidateProps.project.repo,
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
      repoBranch: ValidateProps.project.repoBranch,
      pullRequestPreviewsEnabled:
        ValidateProps.project.pullRequestPreviewsEnabled,
    },
    additionalProperties: true,
  }),
  async (req, res) => {
    if (!req.user) {
      return res.status(401).end();
    }

    if (!req.user.emailVerified && process.env.NEXT_PUBLIC_FULL == "true") {
      return res.status(401).json({ message: "You must verify your email" });
    }

    const db = await getMongoDb();

    let {
      ghToken,
      railwayApiToken,
      userId,
      name,
      repo,
      visibility,
      platform,
      lang,
      packageManager,
      hostService,
      botToken,
      botAppToken,
      botSecretToken,
      repoBranch,
      pullRequestPreviewsEnabled,
    } = req.body;

    let prj = {
      name,
      repo,
      botToken,
      botAppToken,
      botSecretToken,
      platform,
      lang,
      packageManager,
      visibility,
      hostService,
      railwayProjectId: "",
      railwayServiceId: "",
      railwayEnvId: "",
      renderServiceId: "",
      icon: "",
      buildCommand: "",
      startCommand: "",
      rootDirectory: "",
      repoBranch,
      pullRequestPreviewsEnabled,
    };

    const { payload: ghApiToken } = await jwtDecrypt(ghToken, BW_SECRET_KEY);

    if (hostService == "railway") {
      const { payload: rwApiToken } = await jwtDecrypt(
        railwayApiToken,
        BW_SECRET_KEY
      );

      const createRailwayProject = await fetcher(
        "https://backboard.railway.app/graphql/v2",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${rwApiToken.data}`,
          },
          body: JSON.stringify({
            operationName: "projectCreate",
            query: `mutation projectCreate { projectCreate(input: { name: "${name}" }) { id }}`,
          }),
        }
      );

      const createService = await fetcher(
        "https://backboard.railway.app/graphql/v2",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${rwApiToken.data}`,
          },
          body: JSON.stringify({
            operationName: "serviceCreate",
            query: `mutation serviceCreate { serviceCreate(input: { name: "${
              name + "-main"
            }", projectId: "${
              createRailwayProject.data.projectCreate.id
            }" }) { id }}`,
          }),
        }
      );

      const railwayProjectId = await new EncryptJWT({
        data: createRailwayProject.data.projectCreate.id,
      })
        .setProtectedHeader({ alg: "dir", enc: "A128CBC-HS256" })
        .encrypt(BW_SECRET_KEY);

      const railwayServiceId = await new EncryptJWT({
        data: createService.data.serviceCreate.id,
      })
        .setProtectedHeader({ alg: "dir", enc: "A128CBC-HS256" })
        .encrypt(BW_SECRET_KEY);

      prj["railwayProjectId"] = railwayProjectId;
      prj["railwayServiceId"] = railwayServiceId;
    }

    const project: any = await insertProject(db, userId, prj);

    const octokit = new Octokit({
      auth: ghApiToken.data,
    });

    const ghu = await (await octokit.request("GET /user", {})).data;

    await octokit.request("POST /user/repos", {
      name,
      description: `My Awesome ${platform} botway bot.`,
      private: visibility != "public",
    });

    exec(
      `create-botway-bot ${stringify(name)} ${stringify(platform)} ${stringify(
        lang
      )} ${stringify(packageManager)} ${stringify(hostService)} ${stringify(
        ghApiToken.data
      )} ${stringify(ghu.login)} ${stringify(ghu.email)}`
    )
      .on("error", (e) => {
        return res.json({ e });
      })
      .on("message", (m) => {
        console.log(m);
      });

    return res.json({ project });
  }
);

export default handler;
