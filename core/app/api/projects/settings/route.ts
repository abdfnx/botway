import { NextResponse } from "next/server";
import { jwtDecrypt } from "jose";
import { BW_SECRET_KEY } from "@/tools/tokens";
import { fetcher } from "@/tools/fetch";
import createClient from "@/supabase/server";
import { Octokit } from "octokit";

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

  const { payload: githubApiToken } = await jwtDecrypt(
    user?.user_metadata["githubApiToken"],
    BW_SECRET_KEY,
  );

  const { payload: railwayApiToken } = await jwtDecrypt(
    user?.user_metadata["railwayApiToken"],
    BW_SECRET_KEY,
  );

  const { payload: railwayProjectId } = await jwtDecrypt(
    body.railwayProjectId,
    BW_SECRET_KEY,
  );

  const { payload: railwayServiceId } = await jwtDecrypt(
    body.railwayServiceId,
    BW_SECRET_KEY,
  );

  const octokit = new Octokit({
    auth: githubApiToken.data,
  });

  const ghu = await (await octokit.request("GET /user", {})).data;

  if (!body.repo.toString().includes(ghu.login))
    NextResponse.json({ error: `Repo owner must be ${ghu.login}` });

  const iconBody = body.icon != "" ? `icon: "${body.icon}"` : "";

  const repoBody = body.repo ? `source: { repo: "${body.repo}" }` : "";

  const buildCommandBody = body.buildCmd
    ? `buildCommand: "${body.buildCmd}"`
    : "";

  const rootDirectoryBody = body.rootDir
    ? `rootDirectory: "${body.rootDir}"`
    : "";

  const startCommandBody = body.startCmd
    ? `startCommand: "${body.startCmd}"`
    : "";

  const query = `
    mutation settingsUpdate {
      projectUpdate(
        id: "${railwayProjectId.data}",
        input: {
          name: "${body.name}"
        }) {
          id
        }

      serviceUpdate(
        id: "${railwayServiceId.data}",
        input: { 
          ${iconBody} 
        }) {
          id
        }

      serviceInstanceUpdate(
        serviceId: "${railwayServiceId.data}",
        input: {
          ${buildCommandBody},
          ${startCommandBody},
          ${rootDirectoryBody},
          ${repoBody}
        }
      )
    }
  `;

  const editBot = await fetcher("https://backboard.railway.app/graphql/v2", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${railwayApiToken.data}`,
    },
    body: JSON.stringify({
      operationName: "settingsUpdate",
      query,
    }),
  });

  if (editBot.errors) {
    return NextResponse.json({ error: editBot.errors[0].message });
  }

  const { error } = await supabase
    .from("projects")
    .update({
      name: body.name,
      icon: body.icon,
      repo: body.repo,
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
