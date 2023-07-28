import { NextResponse } from "next/server";
import { fetcher } from "@/tools/fetch";
import createClient from "@/supabase/server";

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

  const ai = await fetcher(`https://ai-botway.hop.sh/api/new-bot`, {
    method: "POST",
    body: JSON.stringify({
      prompt: body.prompt,
    }),
  });

  return NextResponse.json({ message: "Success", answer: ai.data });
}
