import { NextResponse } from "next/server";
import { fetcher } from "@/tools/fetch";
import { jwtDecrypt } from "jose";
import { BW_SECRET_KEY } from "@/tools/tokens";
import createClient from "@/supabase/server";
import { faker } from "@faker-js/faker";

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

  const { data, error } = await supabase
    .from("projects")
    .select("zeabur_project_id, zeabur_env_id")
    .eq("id", body.projectId)
    .single();

  if (error) {
    return NextResponse.json({ error });
  }

  const { payload: zeaburProjectToken } = await jwtDecrypt(
    data?.zeabur_project_id,
    BW_SECRET_KEY,
  );

  const { payload: zeaburEnvToken } = await jwtDecrypt(
    data?.zeabur_env_id,
    BW_SECRET_KEY,
  );

  const query = `
    mutation {
      createServiceFromMarketplace(
        projectID: "${zeaburProjectToken.data}"
        itemCode: "${body.slug}"
        name: "${body.name}"
      ) {
        _id
      }
    }
  `;

  const deploy = await fetcher("https://gateway.zeabur.com/graphql", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${zeaburApiToken.data}`,
    },
    body: JSON.stringify({
      query,
    }),
  });

  if (deploy.errors) {
    return NextResponse.json({ message: deploy.errors[0].message });
  }

  if (!body.is_plugin) {
    const id = deploy.data.createServiceFromMarketplace._id;

    const query = `
      mutation {
        addDomain(
          serviceID: "${id}"
          environmentID: "${zeaburEnvToken.data}"
          domain: "${faker.word.sample() + "-" + faker.word.sample()}"
          isGenerated: true
        ) {
          _id
          domain
        }
      }
    `;

    const addDomain = await fetcher("https://gateway.zeabur.com/graphql", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${zeaburApiToken.data}`,
      },
      body: JSON.stringify({
        query,
      }),
    });

    if (addDomain.errors) {
      return NextResponse.json({ message: addDomain.errors[0].message });
    }
  }

  return NextResponse.json({ message: "Success" });
}
