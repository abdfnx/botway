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
    .select("railway_project_id")
    .eq("id", body.projectId)
    .single();

  if (error) {
    return NextResponse.json({ error });
  }

  const { payload: railwayProjectId } = await jwtDecrypt(
    data.railway_project_id,
    BW_SECRET_KEY
  );

  const { payload: railwayApiToken } = await jwtDecrypt(
    user?.user_metadata["railwayApiToken"],
    BW_SECRET_KEY
  );

  const projectData = await fetcher(
    "https://backboard.railway.app/graphql/v2",
    {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${railwayApiToken.data}`,
      },
      body: JSON.stringify({
        query: `
          query {
            project(id: "${railwayProjectId.data}") {
              services {
                edges {
                  node {
                    id
                    name
                  }
                }
              }

              plugins {
                edges {
                  node {
                    id
                    friendlyName
                    name
                  }
                }
              }

              volumes {
                edges {
                  node {
                    id
                    name
                  }
                }
              }
            }
          }
      `,
      }),
    }
  );

  if (projectData.errors) {
    return NextResponse.json({ error: projectData.errors[0].message });
  }

  return NextResponse.json({
    message: "Success",
    services: projectData.data.project.services.edges,
    plugins: projectData.data.project.plugins.edges,
    volumes: projectData.data.project.volumes.edges,
  });
}
