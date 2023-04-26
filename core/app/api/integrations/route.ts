import { NextResponse } from "next/server";
import { fetcher } from "@/tools/fetch";

export const revalidate = 0;

export async function GET(_: Request) {
  const data = await fetcher(`https://cdn-botway.deno.dev/integrations`, {
    method: "GET",
  });

  const intx = data.sort((a: any, b: any) => {
    return new Date(a.created_at).getTime() - new Date(b.created_at).getTime();
  });

  return NextResponse.json(intx);
}
