import { createClient } from "supabase";

const options = {
  schema: "public",
  autoRefreshToken: true,
  persistSession: true,
  detectSessionInUrl: true,
};

export const supabase = createClient(
  Deno.env.get("NEXT_PUBLIC_SUPABASE_URL"),
  Deno.env.get("NEXT_PUBLIC_SUPABASE_ANON_KEY"),
  options,
);
