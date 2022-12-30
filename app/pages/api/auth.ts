import { passport } from "@/api/auth";
import { auths } from "@/api/middlewares";
import { ncOpts } from "@/api/nc";
import nc from "next-connect";

const handler = nc(ncOpts);

handler.use(...auths);

handler.post(passport.authenticate("local"), (req: any, res: any) => {
  res.json({ user: req.user });
});

handler.delete(async (req: any, res: any) => {
  await req.session.destroy();

  res.status(204).end();
});

export default handler;
