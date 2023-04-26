import { Handlers, HandlerContext } from "$fresh/server.ts";
import { add } from "../data.ts";

export const handler: Handlers = {
  async GET(_, ctx: HandlerContext) {
    add();

    return new Response(JSON.stringify({ message: "OK 👍" }));
  },
};
