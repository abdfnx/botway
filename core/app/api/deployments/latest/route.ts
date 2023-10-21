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
    .select("zeabur_service_id, zeabur_env_id")
    .eq("id", id)
    .single();

  if (error) {
    return NextResponse.json({ error });
  }

  const { payload: zeaburApiToken } = await jwtDecrypt(
    user?.user_metadata["zeaburApiToken"],
    BW_SECRET_KEY,
  );

  const { payload: zeaburServiceId } = await jwtDecrypt(
    data?.zeabur_service_id,
    BW_SECRET_KEY,
  );

  const { payload: zeaburEnvId } = await jwtDecrypt(
    data?.zeabur_env_id,
    BW_SECRET_KEY,
  );

  const deployments = await fetcher("https://gateway.zeabur.com/graphql", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${zeaburApiToken.data}`,
    },
    body: JSON.stringify({
      query: `
        query {
          deployments(
            serviceID: "${zeaburServiceId.data}"
            environmentID: "${zeaburEnvId.data}"
          ) {
            edges {
              node {
                _id
              }
            }
          }
        }
      `,
    }),
  });

  if (deployments.errors) {
    return NextResponse.json({ error: deployments.errors[0].message });
  }

  return NextResponse.json(deployments.data.deployments.edges[0].node._id);
}
