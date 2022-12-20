import { ValidateProps } from "@/api/constants";
import { updateUserPasswordByOldPassword } from "@/api/db";
import { auths, validateBody } from "@/api/middlewares";
import { getMongoDb } from "@/api/mongodb";
import { ncOpts } from "@/api/nc";
import nc from "next-connect";
import { NextApiRequest, NextApiResponse } from "next";

const handler = nc<NextApiRequest, NextApiResponse>(ncOpts);

handler.use(...auths);

handler.put(
  validateBody({
    type: "object",
    properties: {
      oldPassword: ValidateProps.user.password,
      newPassword: ValidateProps.user.password,
    },
    required: ["oldPassword", "newPassword"],
    additionalProperties: false,
  }),
  async (req: any, res: any) => {
    if (!req.user) {
      res.json(401).end();

      return;
    }

    const db = await getMongoDb();

    const { oldPassword, newPassword } = req.body;

    const success = await updateUserPasswordByOldPassword(
      db,
      req.user._id,
      oldPassword,
      newPassword
    );

    if (!success) {
      res.status(401).json({
        error: { message: "The old password you entered is incorrect." },
      });

      return;
    }

    res.status(204).end();
  }
);

export default handler;
