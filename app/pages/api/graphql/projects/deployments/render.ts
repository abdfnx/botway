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

  const { renderApiToken, renderServiceId } = req.body;

  const deployments = await fetcher(
    `https://api.render.com/v1/services/${renderServiceId}/deploys`,
    {
      method: "GET",
      headers: {
        Accept: "application/json",
        Authorization: `Bearer ${renderApiToken}`,
      },
    }
  );

  const dy = deployments.sort((a: any, b: any) => {
    return (
      new Date(b.deploy.updatedAt).getTime() -
      new Date(a.deploy.updatedAt).getTime()
    );
  });

  res.json(dy);
});

export const config = {
  api: {
    bodyParser: false,
  },
};

export default handler;
