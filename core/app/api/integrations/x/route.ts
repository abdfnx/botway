import { NextResponse } from "next/server";
import { fetcher } from "@/tools/fetch";

export const revalidate = 0;

export async function GET(request: Request) {
  const { searchParams } = new URL(request.url);

  const slug = searchParams.get("slug");

  const intx = await fetcher(
    `https://cdn-botway.deno.dev/integrations/${slug}`,
    {
      method: "GET",
    },
  );

  return NextResponse.json(intx);
}
