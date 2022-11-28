import { nanoid } from "nanoid";

export function findTokenByIdAndType(db: any, id: any, type: any) {
  return db.collection("tokens").findOne({
    _id: id,
    type,
  });
}

export function findAndDeleteTokenByIdAndType(db: any, id: any, type: any) {
  return db
    .collection("tokens")
    .findOneAndDelete({ _id: id, type })
    .then(({ value }: any) => value);
}

export async function createToken(db: any, { creatorId, type, expireAt }: any) {
  const securedTokenId = nanoid(32);

  const token = {
    _id: securedTokenId,
    creatorId,
    type,
    expireAt,
  };

  await db.collection("tokens").insertOne(token);

  return token;
}
