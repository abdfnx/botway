import { supabase } from "./supabase/config.ts";

const folders = ["icons", "images", "fonts", "simple", "screenshots"];

export const add = () => {
  folders.forEach(async (v) => {
    await Deno.mkdir(`./static/${v}`, { recursive: true });

    const { data, error } = await supabase.storage.from("cdn").list(v, {
      limit: 10000,
      offset: 0,
      sortBy: { column: "name", order: "asc" },
    });

    if (error) {
      console.log(JSON.stringify({ status: 404, error }));
    }

    data.map(async (d) => {
      const response = await fetch(
        `${Deno.env.get(
          "NEXT_PUBLIC_SUPABASE_URL",
        )}/storage/v1/object/public/cdn/${v}/${d.name}`,
      );

      if (!response.ok) throw new Error("Response not OK");

      const r = response.body?.getReader;

      const buf = new Uint8Array(await response.arrayBuffer());

      await Deno.writeFile(`./static/${v}/${d.name}`, buf);
    });
  });
};

if (Deno.args[0] === "start") {
  add();
}
