import { supabase } from "../../supabase/config.ts";
import { Handlers } from "$fresh/server.ts";

export const handler: Handlers = {
  async GET(_, ctx) {
    const { data, error } = await supabase.from("main").select("*");

    if (error) {
      return new Response(JSON.stringify({ status: 404, error }));
    }

    return new Response(JSON.stringify(data));
  },
};
