import { supabase } from "../../supabase/config.ts";
import { Handlers, HandlerContext } from "$fresh/server.ts";

export const handler: Handlers = {
  async GET(_, ctx: HandlerContext) {
    const { folder } = ctx.params;

    const { data, error } = await supabase.storage.from("cdn").list(folder, {
      limit: 10000,
      offset: 0,
      sortBy: { column: "name", order: "asc" },
    });

    if (error) {
      return new Response(JSON.stringify({ status: 404, error }));
    }

    data.map(async (d) => {
      const response = await fetch(
        `${Deno.env.get(
          "NEXT_PUBLIC_SUPABASE_URL"
        )}/storage/v1/object/public/cdn/${folder}/${d.name}`
      );

      if (!response.ok) throw new Error("Response not OK");

      const r = response.body?.getReader;

      const buf = new Uint8Array(await response.arrayBuffer());

      await Deno.writeFile(d.name, buf);
    });

    if (error) {
      return new Response(`{"error": "${error}"}`);
    }

    return new Response(`{"folder": "${folder}"}`);
  },
};
