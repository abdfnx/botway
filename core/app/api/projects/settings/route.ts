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

  const { data, error: pError } = await supabase
    .from("projects")
    .select("zeabur_project_id, zeabur_service_id, name")
    .eq("id", body.projectId)
    .single();

  if (pError) {
    return NextResponse.json({ error: pError });
  }

  const { payload: zeaburProjectId } = await jwtDecrypt(
    data.zeabur_project_id,
    BW_SECRET_KEY,
  );

  const { payload: zeaburServiceId } = await jwtDecrypt(
    data.zeabur_service_id,
    BW_SECRET_KEY,
  );

  const nameBody = `
    mutation {
      renameProject(_id: "${zeaburProjectId.data}", name: "${body.name}")
    }
  `;

  const iconBody = body.icon
    ? `updateProjectIcon(projectID: "${zeaburProjectId.data}", iconURL: "${body.icon}")`
    : "";

  const buildCommandBody = body.buildCmd
    ? `updateCustomBuildCommand(serviceID: "${zeaburServiceId.data}", customBuildCommand: "${body.buildCmd}")`
    : "";

  const startCommandBody = body.startCmd
    ? `updateCustomStartCommand(serviceID: "${zeaburServiceId.data}", customStartCommand: "${body.startCmd}")`
    : "";

  const rootDirectoryBody = body.rootDir
    ? `updateRootDirectory(serviceID: "${zeaburServiceId.data}", rootDirectory: "${body.rootDir}")`
    : "";

  const query = `
    mutation {
      ${iconBody}
      ${buildCommandBody}
      ${startCommandBody}
      ${rootDirectoryBody}
    }
  `;

  const editBot = await fetcher("https://gateway.zeabur.com/graphql", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${zeaburApiToken.data}`,
    },
    body: JSON.stringify({
      query,
    }),
  });

  if (editBot.errors) {
    return NextResponse.json({ error: editBot.errors[0].message });
  }

  if (body.name != data.name) {
    const updateName = await fetcher("https://gateway.zeabur.com/graphql", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${zeaburApiToken.data}`,
      },
      body: JSON.stringify({
        query: nameBody,
      }),
    });

    if (updateName.errors) {
      return NextResponse.json({ error: updateName.errors[0].message });
    }
  }

  const { error } = await supabase
    .from("projects")
    .update({
      name: body.name,
      icon: body.icon,
      build_command: body.buildCmd,
      start_command: body.startCmd,
      root_directory: body.rootDir,
    })
    .eq("id", body.projectId);

  if (error) {
    return NextResponse.json({ error });
  }

  return NextResponse.json({ message: "Success" });
}
