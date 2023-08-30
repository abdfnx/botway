import { NextResponse } from "next/server";
import { jwtDecrypt } from "jose";
import { BW_SECRET_KEY } from "@/tools/tokens";
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

  const { payload: zeaburApiToken } = await jwtDecrypt(
    user?.user_metadata["zeaburApiToken"],
    BW_SECRET_KEY,
  );

  const { payload: zeaburProjectId } = await jwtDecrypt(
    body.zeaburProjectId,
    BW_SECRET_KEY,
  );

  const deleteBot = await fetcher("https://gateway.zeabur.com/graphql", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${zeaburApiToken.data}`,
    },
    body: JSON.stringify({
      query: `
        mutation {
          deleteProject(_id: "${zeaburProjectId.data}")
        }
      `,
    }),
  });

  if (deleteBot.errors) {
    return NextResponse.json({ error: deleteBot.errors[0].message });
  }

  const { error } = await supabase
    .from("projects")
    .delete()
    .eq("id", body.projectId);

  if (error) {
    return NextResponse.json({ error });
  }

  return NextResponse.json({ message: "Success" });
}
