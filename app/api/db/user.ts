import bcrypt from "bcryptjs";
import { ObjectId } from "mongodb";
import normalizeEmail from "validator/lib/normalizeEmail";

export async function findUserWithEmailAndPassword(
  db: any,
  email: any,
  password: any
) {
  email = normalizeEmail(email);

  const user = await db.collection("users").findOne({ email });

  if (user && (await bcrypt.compare(password, user.password))) {
    return { ...user, password: undefined }; // filtered out password
  }

  return null;
}

export async function findUserForAuth(db: any, userId: any) {
  return db
    .collection("users")
    .findOne({ _id: new ObjectId(userId) }, { projection: { password: 0 } })
    .then((user: any) => user || null);
}

export async function findUserById(db: any, userId: any) {
  return db
    .collection("users")
    .findOne({ _id: new ObjectId(userId) }, { projection: dbProjectionUsers() })
    .then((user: any) => user || null);
}

export async function findUserByUsername(db: any, username: any) {
  return db
    .collection("users")
    .findOne({ username }, { projection: dbProjectionUsers() })
    .then((user: any) => user || null);
}

export async function findUserByEmail(db: any, email: any) {
  email = normalizeEmail(email);

  return db
    .collection("users")
    .findOne({ email }, { projection: dbProjectionUsers() })
    .then((user: any) => user || null);
}

export async function updateUserById(db: any, id: any, data: any) {
  return db
    .collection("users")
    .findOneAndUpdate(
      { _id: new ObjectId(id) },
      { $set: data },
      { returnDocument: "after", projection: { password: 0 } }
    )
    .then(({ value }: any) => value);
}

export async function insertUser(
  db: any,
  {
    email,
    originalPassword,
    name,
    username,
    isAdmin,
    githubApiToken,
    railwayApiToken,
    renderApiToken,
    renderUserEmail,
  }: any
) {
  let user: any = {
    emailVerified: false,
    email,
    name,
    username,
    isAdmin,
    githubApiToken,
    railwayApiToken,
    renderApiToken,
    renderUserEmail,
  };

  const password = await bcrypt.hash(originalPassword, 10);

  const { insertedId } = await db
    .collection("users")
    .insertOne({ ...user, password });

  user._id = insertedId;

  return user;
}

export async function updateUserPasswordByOldPassword(
  db: any,
  id: any,
  oldPassword: any,
  newPassword: any
) {
  const user = await db.collection("users").findOne(new ObjectId(id));

  if (!user) return false;

  const matched = await bcrypt.compare(oldPassword, user.password);

  if (!matched) return false;

  const password = await bcrypt.hash(newPassword, 10);

  await db
    .collection("users")
    .updateOne({ _id: new ObjectId(id) }, { $set: { password } });

  return true;
}

export async function UNSAFE_updateUserPassword(
  db: any,
  id: any,
  newPassword: any
) {
  const password = await bcrypt.hash(newPassword, 10);

  await db
    .collection("users")
    .updateOne({ _id: new ObjectId(id) }, { $set: { password } });
}

export function dbProjectionUsers(prefix = "") {
  return {
    [`${prefix}password`]: 0,
    [`${prefix}email`]: 0,
    [`${prefix}emailVerified`]: 0,
  };
}
