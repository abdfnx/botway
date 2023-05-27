import { createServerComponentClient } from "@supabase/auth-helpers-nextjs";
import { cookies, headers } from "next/headers";

export default () =>
  createServerComponentClient({
    headers,
    cookies,
  });
