import { ValidateProps } from "@/api/constants";
import { insertProject } from "@/api/db";
import { auths, validateBody } from "@/api/middlewares";
import { getMongoDb } from "@/api/mongodb";
import { ncOpts } from "@/api/nc";
import { fetcher } from "@/lib/fetch";
import nc from "next-connect";
import { Octokit } from "octokit";

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

    const db = await getMongoDb();

    const {
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

    if (hostService == "railway") {
      const createRailwayProject = await fetcher(
        "https://backboard.railway.app/graphql/v2",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${railwayApiToken}`,
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
            Authorization: `Bearer ${railwayApiToken}`,
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

      prj["railwayProjectId"] = createRailwayProject.data.projectCreate.id;
      prj["railwayServiceId"] = createService.data.serviceCreate.id;
    }

    const project = await insertProject(db, userId, prj);

    const octokit = new Octokit({
      auth: ghToken,
    });

    const ghu = await (await octokit.request("GET /user", {})).data;

    await octokit.request("POST /user/repos", {
      name,
      description: `My Awesome ${platform} botway bot.`,
      private: visibility != "public",
    });

    await fetcher("https://create-botway-bot.up.railway.app/create", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: req.body.ghToken,
      },
      body: JSON.stringify({
        name,
        visibility,
        platform,
        lang,
        packageManager,
        hostService,
        username: ghu.login,
        email: ghu.email,
      }),
    });

    return res.json({ project });
  }
);

export default handler;
