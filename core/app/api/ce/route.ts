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

  const { payload: railwayApiToken } = await jwtDecrypt(
    user?.user_metadata["railwayApiToken"],
    BW_SECRET_KEY
  );

  const { payload: railwayProjectId } = await jwtDecrypt(
    body.railwayProjectId,
    BW_SECRET_KEY
  );

  const query = `
    query {
      project(id: "${railwayProjectId.data}") {
        services {
          edges {
            node {
              id
              serviceInstances {
                edges {
                  node {
                    domains {
                      customDomains {
                        domain
                      }

                      serviceDomains {
                        domain
                      }
                    }

                    source {
                      repo
                      template {
                        serviceSource
                      }
                    }
                  }
                }
              }
            }
          }
        }
      }
    }
  `;

  const check = await fetcher("https://backboard.railway.app/graphql/v2", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${railwayApiToken.data}`,
    },
    body: JSON.stringify({
      query,
    }),
  });

  if (check.errors) {
    console.log(check.errors);

    return NextResponse.json({ message: check.errors[0].message });
  }

  const domainNode = check.data.project.services.edges.find((srv: any) =>
    srv.node.serviceInstances.edges.find(
      (si: any) =>
        si.node.source.template?.serviceSource ===
        "https://github.com/botwayorg/ce"
    )
  ).node.serviceInstances.edges[0].node.domains;

  let domain;

  if (domainNode.customDomains.length != 0) {
    domain = domainNode.customDomains[0].domain;
  } else {
    domain = domainNode.serviceDomains[0].domain;
  }

  return NextResponse.json({ message: "Success", domain });
}
