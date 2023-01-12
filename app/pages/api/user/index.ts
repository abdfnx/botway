import { ValidateProps } from "@/api/constants";
import { findUserByUsername, updateUserById } from "@/api/db";
import { auths, validateBody } from "@/api/middlewares";
import { getMongoDb } from "@/api/mongodb";
import { ncOpts } from "@/api/nc";
import { slugger } from "@/lib/user";
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
      username: ValidateProps.user.username,
      name: ValidateProps.user.name,
      githubApiToken: ValidateProps.user.githubApiToken,
      railwayApiToken: ValidateProps.user.railwayApiToken,
      renderApiToken: ValidateProps.user.renderApiToken,
    },
    additionalProperties: true,
  }),
  async (req, res) => {
    if (!req.user) {
      req.status(401).end();

      return;
    }

    const db = await getMongoDb();

    const { name, githubApiToken, railwayApiToken, renderApiToken } = req.body;

    let username;

    if (req.body.username) {
      username = slugger(req.body.username);

      if (
        username !== req.user.username &&
        (await findUserByUsername(db, username))
      ) {
        res
          .status(403)
          .json({ error: { message: "The username has already been taken." } });
        return;
      }
    }

    const user = await updateUserById(db, req.user._id, {
      ...(username && { username }),
      ...(name && { name }),
      githubApiToken,
      railwayApiToken,
      renderApiToken,
    });

    res.json({ user });
  }
);

export const config = {
  api: {
    bodyParser: false,
  },
};

export default handler;
