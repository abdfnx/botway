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

  const { payload: railwayApiToken } = await jwtDecrypt(
    user?.user_metadata["railwayApiToken"],
    BW_SECRET_KEY
  );

  const { data, error } = await supabase
    .from("projects")
    .select("railway_project_id, railway_service_id")
    .eq("id", body.projectId)
    .single();

  if (error) {
    return NextResponse.json({ error });
  }

  const { payload: railwayProjectId } = await jwtDecrypt(
    data.railway_project_id,
    BW_SECRET_KEY
  );

  const { payload: railwayServiceId } = await jwtDecrypt(
    data.railway_service_id,
    BW_SECRET_KEY
  );

  const { payload: value } = await jwtDecrypt(body.value, BW_SECRET_KEY);

  const getEnvId = await fetcher("https://backboard.railway.app/graphql/v2", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${railwayApiToken.data}`,
    },
    body: JSON.stringify({
      query: `
        query {
          project(id: "${railwayProjectId.data}") {
            environments {
              edges {
                node {
                  name,
                  id
                }
              }
            }
          }
        }
      `,
    }),
  });

  if (getEnvId.errors) {
    return NextResponse.json({ error: getEnvId.errors[0].message });
  }

  const envId = getEnvId.data.project.environments.edges.find(
    (env: any) => env.node.name === "production"
  ).node.id;

  const updateVar = await fetcher("https://backboard.railway.app/graphql/v2", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${railwayApiToken.data}`,
    },
    body: JSON.stringify({
      query: `
        mutation {
          variableUpsert(input: {
            environmentId: "${envId}"
            name: "${body.key}"
            projectId: "${railwayProjectId.data}"
            serviceId: "${railwayServiceId.data}"
            value: "${value.data}"
          })
        }
      `,
    }),
  });

  if (updateVar.errors) {
    return NextResponse.json({ error: updateVar.errors[0].message });
  }

  return NextResponse.json({
    message: "Success",
  });
}
