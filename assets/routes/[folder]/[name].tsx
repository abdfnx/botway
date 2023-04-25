import { Handlers, HandlerContext } from "$fresh/server.ts";

export const handler: Handlers = {
  async GET(_, ctx: HandlerContext) {
    const { name } = ctx.params;

    const imageBuf = await Deno.readFile(`${Deno.cwd()}/${name}`);

    const resp = await new Response(imageBuf);

    if (name.includes("svg")) {
      resp.headers.set("Content-Type", "image/svg+xml");
    } else if (name.includes("png")) {
      resp.headers.set("Content-Type", "image/png");
    } else {
      resp.headers.set("Content-Type", "application/octet-stream");
    }

    return resp;
  },
};
