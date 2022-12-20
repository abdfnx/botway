import { findUserById } from "@/api/db";
import { getMongoDb } from "@/api/mongodb";
import { ncOpts } from "@/api/nc";
import nc from "next-connect";
import { NextApiRequest, NextApiResponse } from "next";

const handler = nc<NextApiRequest, NextApiResponse>(ncOpts);

handler.get(async (req: any, res: any) => {
  const db = await getMongoDb();
  const user = await findUserById(db, req.query.userId);

  res.json({ user });
});

export default handler;
