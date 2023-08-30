import { NextResponse } from "next/server";
import { jwtDecrypt } from "jose";
import { BW_SECRET_KEY } from "@/tools/tokens";
import { fetcher } from "@/tools/fetch";
import createClient from "@/supabase/server";
import { createClient as cc } from "graphql-ws";
import WebSocket from "ws";

const connection = cc({
  url: "wss://gateway.zeabur.com/graphql",
  webSocketImpl: WebSocket,
  connectionParams() {
    return {
      headers: {
        authToken:
          "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ6ZWFidXIiLCJqdGkiOiI2NGM3OTAzMWQxNDlhOTk5MTEzYzYzY2ExNjkwODAwMTc3IiwiaWF0IjoxNjkwODAwMTc3LCJzdWIiOiI2NGM3OTAzMWQxNDlhOTk5MTEzYzYzY2EiLCJmcm9tIjoiemVhYnVyIiwic2NvcGUiOiJhbGwifQ.iW15x2V04lq18zyqXj_JPcmDQskklBj3DDpqJ8o2T48",
      },
    };
  },
});

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
    data.zeabur_service_id,
    BW_SECRET_KEY,
  );

  const { payload: zeaburEnvId } = await jwtDecrypt(
    data.zeabur_env_id,
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
            serviceID: "${zeaburServiceId.data}",
            environmentID: "${zeaburEnvId.data}"
          ) {
            edges {
              node {
                _id
                createdAt
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

  if (deployments.data.deployments.edges.length === 0) {
    return NextResponse.json({ message: "No Logs" });
  } else {
    const dy = deployments.data.deployments.edges.sort((a: any, b: any) => {
      return (
        new Date(b.node.createdAt).getTime() -
        new Date(a.node.createdAt).getTime()
      );
    });

    console.log(dy[0].node._id);

    const SUBSCRIPTION = `
      subscription {
        buildLogReceived(
          deploymentID: "64d7d9fcc640bb43db7dd396"
        ) {
          message
          timestamp
        }
      }
    `;

    let x: any = [];

    connection.subscribe(
      {
        query: SUBSCRIPTION,
      },
      {
        next: (data) => {
          // Handle each received data item
          console.log(data);

          x.push(data);
        },
        error: (error) => {
          // Handle subscription errors
          console.error("Subscription error:", error);
        },
        complete: () => {
          // Handle subscription completion
          console.log("Subscription completed");
        },
      },
    );

    console.log(x);

    return NextResponse.json({
      // logs: logs.data.deploymentLogs,
      dyId: dy[0].node.id,
      message: "Success",
    });
  }
}
