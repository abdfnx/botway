import { ValidateProps } from "@/api/constants";
import { findProjects, insertProject } from "@/api/db";
import { auths, validateBody } from "@/api/middlewares";
import { getMongoDb } from "@/api/mongodb";
import { ncOpts } from "@/api/nc";
import { fetcher } from "@/lib/fetch";
import nc from "next-connect";

const handler = nc(ncOpts);

handler.get(async (req, res) => {
  const db = await getMongoDb();

  const projects = await findProjects(
    db,
    req.query.before ? new Date(req.query.before) : undefined,
    req.query.by
  );

  res.json({ projects });
});

handler.post(
  ...auths,
  validateBody({
    type: "object",
    properties: {
      name: ValidateProps.project.name,
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
      name,
      platform,
      lang,
      packageManager,
      hostService,
      botToken,
      botAppToken,
      botSecretToken,
    } = req.body;

    const rw = await fetcher("https://backboard.railway.app/graphql/v2", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: req.body.apiToken,
      },
      body: JSON.stringify({
        operationName: "projectCreate",
        query: `mutation projectCreate { projectCreate(input: { name: "${req.body.name}" }) { id }}`,
      }),
    });

    const project = await insertProject(db, {
      creatorId: req.user._id,
      name,
      platform,
      lang,
      packageManager,
      hostService,
      botToken,
      botAppToken,
      botSecretToken,
      railwayProjectId: rw.data.projectCreate.id,
      renderProjectId: "",
    });

    return res.json({ project });
  }
);

export default handler;
