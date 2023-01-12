import { ObjectId } from "mongodb";
import { randomUUID } from "crypto";

export async function insertProject(
  db: any,
  userId: any,
  {
    name,
    repo,
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
    renderServiceId,
    icon,
    buildCommand,
    startCommand,
    rootDirectory,
    repoBranch,
    pullRequestPreviewsEnabled,
  }: any
) {
  const project: any = {
    id: randomUUID(),
    name,
    repo,
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
    renderServiceId,
    icon,
    buildCommand,
    startCommand,
    rootDirectory,
    repoBranch,
    pullRequestPreviewsEnabled,
    createdAt: new Date(),
  };

  return db
    .collection("users")
    .findOneAndUpdate(
      { _id: new ObjectId(userId) },
      {
        $push: {
          projects: project,
        },
      },
      { returnDocument: "after", projection: { password: 0 } }
    )
    .then(({ value }: any) => value);
}

export async function deleteProject(
  db: any,
  userId: any,
  id: any,
  { name }: any
) {
  const project: any = {
    id,
    name,
  };

  return db
    .collection("users")
    .findOneAndUpdate(
      { _id: new ObjectId(userId) },
      {
        $pull: {
          projects: project,
        },
      },
      { returnDocument: "after", projection: { password: 0 } }
    )
    .then(({ value }: any) => value);
}

export async function updateProject(db: any, userId: any, id: any, data: any) {
  return db
    .collection("users")
    .findOneAndUpdate(
      { _id: new ObjectId(userId) },
      {
        $set: {
          "projects.$[orderItem]": data,
        },
      },
      {
        returnDocument: "after",
        projection: { password: 0 },
        arrayFilters: [
          {
            "orderItem.id": id,
          },
        ],
      }
    )
    .then(({ value }: any) => value);
}
