import {
  ActivityTypes,
  createBot,
  enableCachePlugin,
  enableCacheSweepers,
  fastFileLoader,
  getAppId,
  getToken,
  startBot,
} from "./deps.ts";
import { logger } from "./src/utils/logger.ts";
import { events } from "./src/events/mod.ts";
import { updateCommands } from "./src/utils/helpers.ts";

const log = logger({ name: "Main" });
const paths = ["./src/events", "./src/commands"];

log.info("Starting Bot, this might take a while...");

await fastFileLoader(paths).catch((err) => {
  log.fatal(`Unable to Import ${paths}`);
  log.fatal(err);

  Deno.exit(1);
});

export const bot = enableCachePlugin(
  createBot({
    token: getToken(),
    botId: getAppId(),
    intents: [],
    events,
  })
);

enableCacheSweepers(bot);

bot.gateway.presence = {
  status: "online",
  activities: [
    {
      name: "Built with Discordeno and Running with Botway",
      type: ActivityTypes.Game,
      createdAt: Date.now(),
    },
  ],
};

await startBot(bot);

await updateCommands(bot);
