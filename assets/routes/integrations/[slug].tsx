import { supabase } from "../../supabase/config.ts";
import { Handlers } from "$fresh/server.ts";

export const handler: Handlers = {
  async GET(_, ctx) {
    const { slug } = ctx.params;

    const { data, error } = await supabase
      .from("main")
      .select("*")
      .eq("slug", slug)
      .single();

    if (error) {
      return new Response(JSON.stringify({ status: 404, error }));
    }

    return new Response(JSON.stringify(data));
  },
};
