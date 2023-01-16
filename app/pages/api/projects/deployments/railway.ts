import { ncOpts } from "@/api/nc";
import nc from "next-connect";
import { fetcher } from "@/lib/fetch";
import { auths } from "@/api/middlewares";
import multer from "multer";
import { BW_SECRET_KEY } from "@/tools/api-tokens";
import { jwtDecrypt } from "jose";

const handler = nc(ncOpts);

handler.use(...auths);

handler.patch(multer({ dest: "/tmp" }).single("data"), async (req, res) => {
  if (!req.user) {
    return req.status(401).end();
  }

  if (!req.user.emailVerified && process.env.NEXT_PUBLIC_FULL == "true") {
    return res.status(401).json({ message: "You must verify your email" });
  }

  let { railwayApiToken, railwayProjectId } = req.body;

  const { payload: rwApiToken } = await jwtDecrypt(
    railwayApiToken,
    BW_SECRET_KEY
  );
  const { payload: rwProjectId } = await jwtDecrypt(
    railwayProjectId,
    BW_SECRET_KEY
  );

  const deployments = await fetcher(
    "https://backboard.railway.app/graphql/v2",
    {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${rwApiToken.data}`,
      },
      body: JSON.stringify({
        query: `query { project(id: "${rwProjectId.data}") { deployments { edges { node { id, createdAt, status, url, meta } } } } }`,
      }),
    }
  );

  const dy = deployments.data.project.deployments.edges.sort(
    (a: any, b: any) => {
      return (
        new Date(b.node.createdAt).getTime() -
        new Date(a.node.createdAt).getTime()
      );
    }
  );

  res.json(dy);
});

export const config = {
  api: {
    bodyParser: false,
  },
};

export default handler;
