import { NextResponse } from "next/server";
import { fetcher } from "@/tools/fetch";
import { jwtDecrypt } from "jose";
import { BW_SECRET_KEY } from "@/tools/tokens";
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
    .select("zeabur_env_id, ce_service_id")
    .eq("id", body.projectId)
    .single();

  if (error) {
    return NextResponse.json({ error });
  }

  const { payload: zeaburApiToken } = await jwtDecrypt(
    user?.user_metadata["zeaburApiToken"],
    BW_SECRET_KEY,
  );

  const { payload: ceServiceId } = await jwtDecrypt(
    data.ce_service_id,
    BW_SECRET_KEY,
  );

  const { payload: zeaburEnvId } = await jwtDecrypt(
    data.zeabur_env_id,
    BW_SECRET_KEY,
  );

  const query = `
    query {
      service(_id: "${ceServiceId.data}") {
        domains(environmentID: "${zeaburEnvId.data}") {
          domain
        }
      }
    }
  `;

  const check = await fetcher("https://gateway.zeabur.com/graphql", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${zeaburApiToken.data}`,
    },
    body: JSON.stringify({
      query,
    }),
  });

  if (check.errors) {
    return NextResponse.json({ message: check.errors[0].message });
  }

  return NextResponse.json({
    message: "Success",
    domain: check.data.service.domains[0].domain,
  });
}
