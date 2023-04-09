import { ValidateProps } from "@/api/constants";
import { findUserByEmail, findUserByUsername, insertUser } from "@/api/db";
import { auths, validateBody } from "@/api/middlewares";
import { getMongoDb } from "@/api/mongodb";
import { ncOpts } from "@/api/nc";
import { slugger } from "@/lib/user";
import nc from "next-connect";
import isEmail from "validator/lib/isEmail";
import normalizeEmail from "validator/lib/normalizeEmail";

const handler = nc(ncOpts);

handler.post(
  validateBody({
    type: "object",
    properties: {
      username: ValidateProps.user.username,
      name: ValidateProps.user.name,
      password: ValidateProps.user.password,
      email: ValidateProps.user.email,
      githubApiToken: ValidateProps.user.githubApiToken,
      railwayApiToken: ValidateProps.user.railwayApiToken,
      renderApiToken: ValidateProps.user.renderApiToken,
    },
    required: ["username", "name", "password", "email"],
    additionalProperties: false,
  }),
  ...auths,
  async (req: any, res: any) => {
    const db = await getMongoDb();

    let {
      username,
      name,
      email,
      password,
      githubApiToken,
      railwayApiToken,
      renderApiToken,
    } = req.body;

    username = slugger(req.body.username);
    email = normalizeEmail(req.body.email);

    if (!isEmail(email)) {
      res
        .status(400)
        .json({ error: { message: "The email you entered is invalid." } });

      return;
    }
    if (await findUserByEmail(db, email)) {
      res
        .status(403)
        .json({ error: { message: "The email has already been used." } });

      return;
    }
    if (await findUserByUsername(db, username)) {
      res
        .status(403)
        .json({ error: { message: "The username has already been taken." } });

      return;
    }

    const user = await insertUser(db, {
      email,
      originalPassword: password,
      name,
      username,
      isAdmin: false,
      githubApiToken,
      railwayApiToken,
      renderApiToken,
    });

    req.logIn(user, (err: any) => {
      if (err) throw err;

      res.status(201).json({
        user,
      });
    });
  }
);

export default handler;
