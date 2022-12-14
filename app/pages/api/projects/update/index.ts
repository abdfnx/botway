import { ValidateProps } from "@/api/constants";
import { updateProject } from "@/api/db";
import { auths, validateBody } from "@/api/middlewares";
import { getMongoDb } from "@/api/mongodb";
import { ncOpts } from "@/api/nc";
import multer from "multer";
import nc from "next-connect";

const handler = nc(ncOpts);

handler.use(...auths);

handler.get(async (req: any, res: any) => {
  if (!req.user) return res.json({ user: null });

  return res.json({ user: req.user });
});

handler.patch(
  multer({ dest: "/tmp" }).single("data"),
  validateBody({
    type: "object",
    properties: {
      name: ValidateProps.project.name,
      platform: ValidateProps.project.platform,
      botToken: ValidateProps.project.botToken,
      botAppToken: ValidateProps.project.botAppToken,
      botSecretToken: ValidateProps.project.botSecretToken,
    },
    additionalProperties: true,
  }),
  async (req, res) => {
    if (!req.user) {
      req.status(401).end();

      return;
    }

    const db = await getMongoDb();

    let { id, name, botToken, platform, botAppToken, botSecretToken } =
      req.body;

    let payload = {
      ...(name && { name }),
      botToken,
    };

    if (platform != "telegram") {
      payload = {
        ...(name && { name }),
        botToken,
        botAppToken,
      };
    } else if (platform == "slack" || platform == "twitch") {
      payload = {
        ...(name && { name }),
        botToken,
        botAppToken,
        botSecretToken,
      };
    }

    const prj = await updateProject(db, id, payload);

    res.json({ prj });
  }
);

export const config = {
  api: {
    bodyParser: false,
  },
};

export default handler;
