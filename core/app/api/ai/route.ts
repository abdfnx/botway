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

  const ai = await fetcher(`https://ai-botway.hop.sh/ai`, {
    method: "POST",
    body: JSON.stringify({
      userId: user?.id.toString(),
      prompt: body.prompt,
    }),
  });

  return NextResponse.json({ message: "Success", answer: ai.data });
}
