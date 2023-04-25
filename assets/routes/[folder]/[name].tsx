import { Handlers, HandlerContext } from "$fresh/server.ts";

export const handler: Handlers = {
  async GET(_, ctx: HandlerContext) {
    const { folder, name } = ctx.params;

    const response = await fetch(
      `${Deno.env.get(
        "NEXT_PUBLIC_SUPABASE_URL"
      )}/storage/v1/object/public/cdn/${folder}/${name}`
    );

    if (!response.ok) throw new Error("Response not OK");

    const r = response.body?.getReader;

    const buf = new Uint8Array(await response.arrayBuffer());

    const resp = await new Response(buf);

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
