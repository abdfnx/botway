import { ncOpts } from "@/api/nc";
import nc from "next-connect";
import { fetcher } from "@/lib/fetch";
import { auths } from "@/api/middlewares";
import multer from "multer";
import { jwtDecrypt } from "jose";
import { BW_SECRET_KEY } from "@/tools/api-tokens";

const handler = nc(ncOpts);

handler.use(...auths);

handler.patch(multer({ dest: "/tmp" }).single("data"), async (req, res) => {
  if (!req.user) {
    req.status(401).end();

    return;
  }

  let { renderApiToken, renderServiceId } = req.body;

  const { payload: rndApiToken } = await jwtDecrypt(
    renderApiToken,
    BW_SECRET_KEY
  );
  const { payload: rndServiceId } = await jwtDecrypt(
    renderServiceId,
    BW_SECRET_KEY
  );

  const deployments = await fetcher(
    `https://api.render.com/v1/services/${rndServiceId.data}/deploys`,
    {
      method: "GET",
      headers: {
        Accept: "application/json",
        Authorization: `Bearer ${rndApiToken.data}`,
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
