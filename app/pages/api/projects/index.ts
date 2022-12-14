import { ValidateProps } from "@/api/constants";
import { findProjects, insertProject } from "@/api/db";
import { auths, validateBody } from "@/api/middlewares";
import { getMongoDb } from "@/api/mongodb";
import { ncOpts } from "@/api/nc";
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
    additionalProperties: false,
  }),
  async (req, res) => {
    if (!req.user) {
      return res.status(401).end();
    }

    const db = await getMongoDb();

    const project = await insertProject(db, {
      creatorId: req.user._id,
      name: req.body.name,
      platform: req.body.platform,
      lang: req.body.lang,
      packageManager: req.body.packageManager,
      hostService: req.body.hostService,
      botToken: req.body.botToken,
      botAppToken: req.body.botAppToken,
      botSecretToken: req.body.botSecretToken,
    });

    return res.json({ project });
  }
);

export default handler;
