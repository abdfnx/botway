import { NextResponse } from "next/server";
import { fetcher } from "@/tools/fetch";
import { jwtDecrypt } from "jose";
import { BW_SECRET_KEY } from "@/tools/tokens";
import createClient from "@/supabase/server";
import { Octokit } from "octokit";

export const revalidate = 0;

export async function POST(request: Request) {
  const body = await request.json();

  const supabase = createClient();

  const {
    data: { user: user },
    error: userError,
  } = await supabase.auth.getUser();

  if (userError) {
    return NextResponse.json({ error: userError });
  }

  const { payload: githubApiToken } = await jwtDecrypt(
    user?.user_metadata["githubApiToken"],
    BW_SECRET_KEY,
  );

  const { payload: railwayApiToken } = await jwtDecrypt(
    user?.user_metadata["railwayApiToken"],
    BW_SECRET_KEY,
  );

  const octokit = new Octokit({
    auth: githubApiToken.data,
  });

  const ghu = await (await octokit.request("GET /user", {})).data;

  return NextResponse.json({ message: "Success" });
}
