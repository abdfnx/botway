import { ncOpts } from "@/api/nc";
import nc from "next-connect";
import { fetcher } from "@/lib/fetch";
import { auths } from "@/api/middlewares";
import multer from "multer";

const handler = nc(ncOpts);

handler.use(...auths);

handler.patch(multer({ dest: "/tmp" }).single("data"), async (req, res) => {
  if (!req.user) {
    req.status(401).end();

    return;
  }

  const { railwayApiToken, railwayProjectId } = req.body;

  const deployments = await fetcher(
    "https://backboard.railway.app/graphql/v2",
    {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${railwayApiToken}`,
      },
      body: JSON.stringify({
        query: `query { project(id: "${railwayProjectId}") { deployments { edges { node { id, createdAt, status, url, meta } } } } }`,
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
