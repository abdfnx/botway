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

export async function findProjects(db: any, before: any, by: any, limit = 10) {
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
      { $limit: limit },
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

export async function insertProject(db: any, { content, creatorId }: any) {
  const project: any = {
    content,
    creatorId,
    createdAt: new Date(),
  };

  const { insertedId } = await db.collection("projects").insertOne(project);

  project._id = insertedId;

  return project;
}
