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
      visibility: ValidateProps.project.visibility,
      platform: ValidateProps.project.platform,
      lang: ValidateProps.project.lang,
      packageManager: ValidateProps.project.packageManager,
      hostService: ValidateProps.project.hostService,
      botToken: ValidateProps.project.botToken,
      botAppToken: ValidateProps.project.botAppToken,
      botSecretToken: ValidateProps.project.botSecretToken,
    },
    additionalProperties: true,
  }),
  async (req, res) => {
    if (!req.user) {
      return res.status(401).end();
    }

    const db = await getMongoDb();

    const {
      userId,
      name,
      visibility,
      platform,
      lang,
      packageManager,
      hostService,
      botToken,
      botAppToken,
      botSecretToken,
    } = req.body;

    const createRailwayProject = await fetcher(
      "https://backboard.railway.app/graphql/v2",
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: req.body.apiToken,
        },
        body: JSON.stringify({
          operationName: "projectCreate",
          query: `mutation projectCreate { projectCreate(input: { name: "${req.body.name}" }) { id }}`,
        }),
      }
    );

    const createService = await fetcher(
      "https://backboard.railway.app/graphql/v2",
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: req.body.apiToken,
        },
        body: JSON.stringify({
          operationName: "serviceCreate",
          query: `mutation serviceCreate { serviceCreate(input: { name: "${
            req.body.name + "-main"
          }", projectId: "${
            createRailwayProject.data.projectCreate.id
          }" }) { id }}`,
        }),
      }
    );

    const project = await insertProject(db, userId, {
      name,
      botToken,
      botAppToken,
      botSecretToken,
      platform,
      lang,
      packageManager,
      visibility,
      hostService,
      railwayProjectId: createRailwayProject.data.projectCreate.id,
      railwayServiceId: createService.data.serviceCreate.id,
      railwayEnvId: "",
      renderProjectId: "",
    });

    const octokit = new Octokit({
      auth: req.body.ghToken,
    });

    const ghu = await (await octokit.request("GET /user", {})).data;

    await octokit.request("POST /user/repos", {
      name,
      description: "My Awesome botway bot.",
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
