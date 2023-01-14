import { deleteProject } from "@/api/db";
import { auths } from "@/api/middlewares";
import { getMongoDb } from "@/api/mongodb";
import { ncOpts } from "@/api/nc";
import { fetcher } from "@/lib/fetch";
import multer from "multer";
import nc from "next-connect";
import { jwtDecrypt } from "jose";
import { BW_SECRET_KEY } from "@/tools/api-tokens";

const handler = nc(ncOpts);

handler.use(...auths);

handler.patch(multer({ dest: "/tmp" }).single("data"), async (req, res) => {
  if (!req.user) {
    return res.status(401).end();
  }

  const db = await getMongoDb();

  let {
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
    const { payload: rwApiToken } = await jwtDecrypt(
      railwayApiToken,
      BW_SECRET_KEY
    );
    const { payload: rwProjectId } = await jwtDecrypt(
      railwayProjectId,
      BW_SECRET_KEY
    );

    await fetcher("https://backboard.railway.app/graphql/v2", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${rwApiToken.data}`,
      },
      body: JSON.stringify({
        operationName: "projectDelete",
        query: `mutation projectDelete { projectDelete(id: "${rwProjectId.data}") }`,
      }),
    });
  } else if (hostService == "render") {
    const { payload: rndApiToken } = await jwtDecrypt(
      renderApiToken,
      BW_SECRET_KEY
    );
    const { payload: rndServiceId } = await jwtDecrypt(
      renderServiceId,
      BW_SECRET_KEY
    );

    await fetcher(`https://api.render.com/v1/services/${rndServiceId.data}`, {
      method: "DELETE",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
        Authorization: `Bearer ${rndApiToken.data}`,
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
