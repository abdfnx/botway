import { ObjectId } from "mongodb";
import { dbProjectionUsers } from "./user";

export async function findProjectById(db: any, id: any) {
  const projects = await db
    .collection("projects")
    .aggregate([
      { $match: { _id: new ObjectId(id) } },
      { $limit: 1 },
      {
        $lookup: {
          from: "users",
          localField: "creatorId",
          foreignField: "_id",
          as: "creator",
        },
      },
      { $unwind: "$creator" },
      { $project: dbProjectionUsers("creator.") },
    ])
    .toArray();

  if (!projects[0]) return null;

  return projects[0];
}

export async function findProjectByName(db: any, name: any) {
  return db
    .collection("projects")
    .findOne({ name })
    .then((prj: any) => prj || null);
}

export async function findProjects(db: any, before: any, by: any) {
  return db
    .collection("projects")
    .aggregate([
      {
        $match: {
          ...(by && { creatorId: new ObjectId(by) }),
          ...(before && { createdAt: { $lt: before } }),
        },
      },
      { $sort: { _id: -1 } },
      {
        $lookup: {
          from: "users",
          localField: "creatorId",
          foreignField: "_id",
          as: "creator",
        },
      },
      { $unwind: "$creator" },
      { $project: dbProjectionUsers("creator.") },
    ])
    .toArray();
}

export async function insertProject(
  db: any,
  {
    name,
    visibility,
    platform,
    lang,
    packageManager,
    hostService,
    creatorId,
    botToken,
    botAppToken,
    botSecretToken,
    railwayProjectId,
    railwayServiceId,
    railwayEnvId,
    renderProjectId,
  }: any
) {
  const project: any = {
    creatorId,
    name,
    visibility,
    platform,
    lang,
    packageManager,
    hostService,
    botToken,
    botAppToken,
    botSecretToken,
    railwayProjectId,
    railwayServiceId,
    railwayEnvId,
    renderProjectId,
    createdAt: new Date(),
  };

  const { insertedId } = await db.collection("projects").insertOne(project);

  project._id = insertedId;

  return project;
}

export async function updateProject(db: any, id: any, data: any) {
  return db
    .collection("projects")
    .findOneAndUpdate(
      { _id: new ObjectId(id) },
      { $set: data },
      { returnDocument: "after" }
    )
    .then(({ value }: any) => value);
}
