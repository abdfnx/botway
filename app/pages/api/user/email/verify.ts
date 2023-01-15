import { createToken } from "@/api/db";
import { sendMail } from "@/api/mail";
import { auths } from "@/api/middlewares";
import { getMongoDb } from "@/api/mongodb";
import { ncOpts } from "@/api/nc";
import nc from "next-connect";
import { EmailTemplate } from "./email-tmpl";

const handler = nc(ncOpts);

handler.use(...auths);

handler.post(async (req, res) => {
  if (!req.user) {
    res.json(401).end();

    return;
  }

  const db = await getMongoDb();

  const token = await createToken(db, {
    creatorId: req.user._id,
    type: "emailVerify",
    expireAt: new Date(Date.now() + 1000 * 60 * 60 * 24),
  });

  await sendMail(
    req.user.email,
    `Botway - Verification Email`,
    EmailTemplate(req.user.name, token._id, req.headers.host, "verify-email")
  );

  res.status(204).end();
});

export default handler;
