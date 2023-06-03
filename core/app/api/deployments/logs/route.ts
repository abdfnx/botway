import { NextResponse } from "next/server";
import { jwtDecrypt } from "jose";
import { BW_SECRET_KEY } from "@/tools/tokens";
import { fetcher } from "@/tools/fetch";
import createClient from "@/supabase/server";

export const revalidate = 0;

export async function GET(request: Request) {
  const { searchParams } = new URL(request.url);

  const id = searchParams.get("id");

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
    .select("railway_service_id")
    .eq("id", id)
    .single();

  if (error) {
    return NextResponse.json({ error });
  }

  const { payload: railwayApiToken } = await jwtDecrypt(
    user?.user_metadata["railwayApiToken"],
    BW_SECRET_KEY
  );

  const { payload: railwayServiceId } = await jwtDecrypt(
    data.railway_service_id,
    BW_SECRET_KEY
  );

  const deployments = await fetcher(
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
            service(id: "${railwayServiceId.data}") {
              deployments {
                edges {
                  node {
                    id,
                    createdAt,
                    status,
                    url,
                    meta
                  }
                }
              }
            }
          }
        `,
      }),
    }
  );

  if (deployments.errors) {
    return NextResponse.json({ error: deployments.errors[0].message });
  }

  if (deployments.data.service.deployments.edges.length === 0) {
    return NextResponse.json({ message: "No Logs" });
  } else {
    const dy = deployments.data.service.deployments.edges.sort(
      (a: any, b: any) => {
        return (
          new Date(b.node.createdAt).getTime() -
          new Date(a.node.createdAt).getTime()
        );
      }
    );

    const logs = await fetcher("https://backboard.railway.app/graphql/v2", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${railwayApiToken.data}`,
      },
      body: JSON.stringify({
        query: `
          query {
            deploymentLogs(limit: 10000, deploymentId: "${dy[0].node.id}") {
              message
              severity
              timestamp
            }
          }
        `,
      }),
    });

    return NextResponse.json({
      logs: logs.data.deploymentLogs,
      dyId: dy[0].node.id,
      message: "Success",
    });
  }
}
