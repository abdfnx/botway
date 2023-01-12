import { deleteProject } from "@/api/db";
import { auths } from "@/api/middlewares";
import { getMongoDb } from "@/api/mongodb";
import { ncOpts } from "@/api/nc";
import { fetcher } from "@/lib/fetch";
import multer from "multer";
import nc from "next-connect";

const handler = nc(ncOpts);

handler.use(...auths);

handler.patch(multer({ dest: "/tmp" }).single("data"), async (req, res) => {
  if (!req.user) {
    return res.status(401).end();
  }

  const db = await getMongoDb();

  const {
    id,
    userId,
    name,
    hostService,
    railwayProjectId,
    railwayApiToken,
    renderServiceId,
    renderApiToken,
  } = req.body;

  if (hostService == "railway") {
    await fetcher("https://backboard.railway.app/graphql/v2", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${railwayApiToken}`,
      },
      body: JSON.stringify({
        operationName: "projectDelete",
        query: `mutation projectDelete { projectDelete(id: "${railwayProjectId}") }`,
      }),
    });
  } else if (hostService == "render") {
    await fetcher(`https://api.render.com/v1/services/${renderServiceId}`, {
      method: "DELETE",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
        Authorization: `Bearer ${renderApiToken}`,
      },
    });
  }

  const project = await deleteProject(db, userId, id, {
    name,
  });

  return res.json({ project });
});

export const config = {
  api: {
    bodyParser: false,
  },
};

export default handler;
