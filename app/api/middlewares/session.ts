import { getMongoClient } from "api/mongodb";
import MongoStore from "connect-mongo";
import nextSession from "next-session";
import { promisifyStore } from "next-session/lib/compat";

const mongoStore = MongoStore.create({
  clientPromise: getMongoClient(),
  stringify: false,
});

const getSession = nextSession({
  store: promisifyStore(mongoStore),
  cookie: {
    httpOnly: true,
    secure: process.env.NODE_ENV === "production",
    maxAge: 2 * 7 * 24 * 60 * 60, // 2 weeks,
    path: "/",
    sameSite: "strict",
  },
  touchAfter: 1 * 7 * 24 * 60 * 60, // 1 week
});

export default async function session(req: any, res: any, next: any) {
  await getSession(req, res);

  next();
}
