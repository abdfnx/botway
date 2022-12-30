import { ObjectId } from "mongodb";
import { randomUUID } from "crypto";

export async function insertProject(
  db: any,
  userId: any,
  {
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
    redisPluginId,
    mysqlPluginId,
    postgresqlPluginId,
    mongodbPluginId,
    renderProjectId,
  }: any
) {
  const project: any = {
    id: randomUUID(),
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
    redisPluginId,
    mysqlPluginId,
    postgresqlPluginId,
    mongodbPluginId,
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
