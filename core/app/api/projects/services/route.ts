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

  const { data, error } = await supabase
    .from("projects")
    .select("zeabur_project_id")
    .eq("id", body.projectId)
    .single();

  if (error) {
    return NextResponse.json({ error });
  }

  const { payload: zeaburApiToken } = await jwtDecrypt(
    user?.user_metadata["zeaburApiToken"],
    BW_SECRET_KEY,
  );

  const { payload: zeaburProjectId } = await jwtDecrypt(
    data.zeabur_project_id,
    BW_SECRET_KEY,
  );

  const projectData = await fetcher("https://gateway.zeabur.com/graphql", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${zeaburApiToken.data}`,
    },
    body: JSON.stringify({
      query: `
        query {
          project(_id: "${zeaburProjectId.data}") {
            services {
              _id
              name
              marketplaceItem {
                code
                iconURL
              }
            }
          }
        }
      `,
    }),
  });

  if (projectData.errors) {
    return NextResponse.json({ error: projectData.errors[0].message });
  }

  return NextResponse.json({
    message: "Success",
    services: projectData.data.project.services,
  });
}
