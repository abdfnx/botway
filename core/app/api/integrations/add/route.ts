import { NextResponse } from "next/server";
import { fetcher } from "@/tools/fetch";
import { jwtDecrypt } from "jose";
import { BW_SECRET_KEY } from "@/tools/tokens";
import createClient from "@/supabase/server";
import { Octokit } from "octokit";
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

  const { payload: githubApiToken } = await jwtDecrypt(
    user?.user_metadata["githubApiToken"],
    BW_SECRET_KEY
  );

  const { payload: railwayApiToken } = await jwtDecrypt(
    user?.user_metadata["railwayApiToken"],
    BW_SECRET_KEY
  );

  const { payload: projectId } = await jwtDecrypt(
    body.projectId,
    BW_SECRET_KEY
  );

  const octokit = new Octokit({
    auth: githubApiToken.data,
  });

  const ghu = await (await octokit.request("GET /user", {})).data;

  let vars;

  if (body.vars) {
    if (body.vars.k) {
      vars = `variables: {${body.vars.k}: "${body.vars.v}" ${body.def_vars}}`;
    } else {
      vars = `variables: {${body.vars.k1}: "${body.vars.v1}" ${body.vars.k2}: "${body.vars.v2}" ${body.def_vars}}`;
    }
  } else {
    vars = `variables: {${body.def_vars}}`;
  }

  const volumes = body.has_volume
    ? `
    volumes: [
      {
        mountPath: "${body.volume_path}"
        projectId: "${projectId.data}"
      }
    ]`
    : "";

  const templateQuery = `
    mutation {
      templateDeploy(input: {
        ${body.plugin ? `plugins: ["${body.plugin}"]` : ""}

        services: [
          {
            hasDomain: true
            isPrivate: true
            owner: "${ghu.login}"
            name: "${
              body.slug + "-" + faker.word.sample() + "-" + faker.word.sample()
            }"
            serviceName: "${
              body.slug + "-" + faker.word.sample() + "-" + faker.word.sample()
            }"
            template: "https://github.com/${body.template_repo}"
            ${vars}
            ${volumes}
          }
        ]

        projectId: "${projectId.data}"
      }) {
        projectId
      }
    }
  `;

  const pluginQuery = `
    mutation {
      pluginCreate(input: {
        name: "${body.slug}"
        projectId: "${projectId.data}"
      }) {
        id
      }
    }
  `;

  const query = body.is_plugin ? pluginQuery : templateQuery;

  const deploy = await fetcher("https://backboard.railway.app/graphql/v2", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${railwayApiToken.data}`,
    },
    body: JSON.stringify({
      query,
    }),
  });

  if (deploy.errors) {
    console.log(deploy.errors);

    return NextResponse.json({ message: deploy.errors[0].message });
  }

  return NextResponse.json({ message: "Success" });
}
