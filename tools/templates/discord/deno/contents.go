package deno

import "fmt"

func DockerfileContent(botName string) string {
	return fmt.Sprintf(`FROM alpine:latest
FROM node:alpine
FROM denoland/deno:alpine
FROM botwayorg/botway:latest

ENV DISCORD_BOT_NAME="%s"
ARG DISCORD_TOKEN
ARG DISCORD_CLIENT_ID

ADD . .

RUN apk update && \
	apk add --no-cache --virtual build-dependencies build-base gcc git ffmpeg curl

# Add packages you want
# RUN apk add PACKAGE_NAME

RUN botway init --docker --name ${DISCORD_BOT_NAME}

USER deno

RUN deno cache deps.ts

EXPOSE 8000

ENTRYPOINT ["deno", "run", "--allow-net", "--allow-write", "--allow-read", "--allow-run", "mod.ts"]`, botName)
}

func ModTsContent() string {
	return fmt.Sprintf(`import { ActivityTypes, createBot, enableCachePlugin, enableCacheSweepers, fastFileLoader, startBot, getToken, getBotId } from "./deps.ts";
import { logger } from "./src/utils/logger.ts";
import { events } from "./src/events/mod.ts";
import { updateCommands } from "./src/utils/helpers.ts";

const log = logger({ name: "Main" });
const paths = ["./src/events", "./src/commands"];

log.info("Starting Bot, this might take a while...");

await fastFileLoader(paths).catch((err) => {
	log.fatal(%s);
	log.fatal(err);

	Deno.exit(1);
});

export const bot = enableCachePlugin(
	createBot({
		token: getToken(),
		botId: getBotId(),
		intents: [],
		events,
	}),
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
`, "`Unable to Import ${paths}`")
}

func DepsTsContent() string {
	return `export * from "https://deno.land/x/discordeno@13.0.0-rc31/mod.ts";
export * from "https://deno.land/x/discordeno@13.0.0-rc31/plugins/mod.ts";
export * from "https://deno.land/std@0.117.0/fmt/colors.ts";
export * from "https://deno.land/x/denobw@0.0.1/mod.ts"`
}

func CommandsModTsContent() string {
	return `import {
	ApplicationCommandOption,
	ApplicationCommandTypes,
	Bot,
	Collection,
	Interaction,
} from "../../deps.ts";

export type subCommand = Omit<Command, "subcommands">;
export type subCommandGroup = {
	name: string;
	subCommands: subCommand[];
};
export interface Command {
	name: string;
	description: string;
	usage?: string[];
	options?: ApplicationCommandOption[];
	type: ApplicationCommandTypes;
	/** Defaults to 'Guild' */
	scope?: "Global" | "Guild";
	execute: (bot: Bot, interaction: Interaction) => unknown;
	subcommands?: Array<subCommandGroup | subCommand>;
}

export const commands = new Collection<string, Command>();

export function createCommand(command: Command) {
	commands.set(command.name, command);
}`
}

func CommandsPingTsContent() string {
	return fmt.Sprintf(`import {
	ApplicationCommandTypes,
	InteractionResponseTypes,
} from "../../deps.ts";
import {
	humanizeMilliseconds,
	snowflakeToTimestamp,
} from "../utils/helpers.ts";
import { createCommand } from "./mod.ts";

createCommand({
	name: "ping",
	description: "Ping the Bot!",
	type: ApplicationCommandTypes.ChatInput,
	scope: "Global",
	execute: async (bot, interaction) => {
		const ping = Date.now() - snowflakeToTimestamp(interaction.id);
		await bot.helpers.sendInteractionResponse(
			interaction.id,
			interaction.token,
			{
				type: InteractionResponseTypes.ChannelMessageWithSource,
				data: {
					content: %s,
				},
			}
		);
	},
});`, "`🏓 Pong! Ping ${ping}ms (${humanizeMilliseconds(ping)})`")
}

func EventsGuildCreateTsContent() string {
	return `import { events } from "./mod.ts";
import { updateGuildCommands } from "../utils/helpers.ts";

events.guildCreate = async (bot, guild) =>
  await updateGuildCommands(bot, guild);`
}

func EventsInteractionCreateTsContent() string {
	return fmt.Sprintf(`import {
	ApplicationCommandOptionTypes,
	bgBlack,
	bgYellow,
	black,
	BotWithCache,
	green,
	Guild,
	red,
	white,
	yellow,
} from "../../deps.ts";
import { events } from "./mod.ts";
import { logger } from "../utils/logger.ts";
import {
	getGuildFromId,
	isSubCommand,
	isSubCommandGroup,
} from "../utils/helpers.ts";
import { Command, commands } from "../commands/mod.ts";

const log = logger({ name: "Event: InteractionCreate" });

events.interactionCreate = async (rawBot, interaction) => {
	const bot = rawBot as BotWithCache;

	if (interaction.data && interaction.id) {
		let guildName = "Direct Message";
		let guild = {} as Guild;

		// Set guild, if there was an error getting the guild, then just say it was a DM. (What else are we going to do?)
		if (interaction.guildId) {
			const guildOrVoid = await getGuildFromId(bot, interaction.guildId).catch(
				(err) => {
				log.error(err);
				}
			);

			if (guildOrVoid) {
				guild = guildOrVoid;
				guildName = guild.name;
			}
		}

		log.info(
			%s
		);

		let command: undefined | Command = interaction.data.name
			? commands.get(interaction.data.name)
			: undefined;
		let commandName = command?.name;

		if (command !== undefined) {
			if (interaction.data.name) {
				if (interaction.data.options?.[0]) {
					const optionType = interaction.data.options[0].type;

					if (optionType === ApplicationCommandOptionTypes.SubCommandGroup) {
						// Check if command has subcommand and handle types
						if (!command.subcommands) return;

						// Try to find the subcommand group
						const subCommandGroup = command.subcommands?.find(
							(command) => command.name == interaction.data?.options?.[0].name
						);

						if (!subCommandGroup) return;

						if (isSubCommand(subCommandGroup)) return;

						// Get name of the command which we are looking for
						const targetCmdName =
							interaction.data.options?.[0].options?.[0].name ||
							interaction.data.options?.[0].options?.[0].name;
						if (!targetCmdName) return;

						// Try to find the command
						command = subCommandGroup.subCommands.find(
							(c) => c.name === targetCmdName
						);

						commandName += %s;
					}

					if (optionType === ApplicationCommandOptionTypes.SubCommandGroup) {
						// Check if command has subcommand and handle types
						if (!command?.subcommands) return;

						// Try to find the command
						const found = command.subcommands.find(
							(command) => command.name == interaction.data?.options?.[0].name
						);

						if (!found) return;

						if (isSubCommandGroup(found)) return;

						command = found;
						commandName += %s;
					}
				}

				try {
					if (command) {
						command.execute(rawBot, interaction);
						log.info(
							%s
						);
					} else {
						throw "";
					}
				} catch (err) {
					log.error(
						%s
					);

					err.length ? log.error(err) : undefined;
				}
			} else {
				log.warn(
					%s
				);
			}
		}
	}
};`, "`[Command: ${bgYellow(black(String(interaction.data.name)))} - ${bgBlack(\nwhite(`Trigger`)\n)}] by ${interaction.user.username}#${\ninteraction.user.discriminator\n} in ${guildName}${\nguildName !== \"Direct Message\" ? ` (${guild.id})` : ``\n}`", "` ${subCommandGroup.name} ${command?.name}`", "` ${command?.name}`", "`[Command: ${bgYellow(\nblack(String(interaction.data.name))\n)} - ${bgBlack(green(`Success`))}] by ${\ninteraction.user.username\n}#${interaction.user.discriminator} in ${guildName}${\nguildName !== \"Direct Message\" ? ` (${guild.id})` : ``\n}`", "`[Command: ${bgYellow(\nblack(String(interaction.data.name))\n)} - ${bgBlack(red(`Error`))}] by ${interaction.user.username}#${\ninteraction.user.discriminator\n} in ${guildName}${\nguildName !== \"Direct Message\" ? ` (${guild.id})` : ``\n}`", "`[Command: ${bgYellow(\nblack(String(interaction.data.name))\n)} - ${bgBlack(yellow(`Not Found`))}] by ${\ninteraction.user.username\n}#${interaction.user.discriminator} in ${guildName}${\nguildName !== \"Direct Message\" ? ` (${guild.id})` : ``\n}`")
}

func EventsModTsContent() string {
	return `import { EventHandlers } from "../../deps.ts";

export const events: Partial<EventHandlers> = {};`
}

func EventsReadyTsContent() string {
	return `import { events } from "./mod.ts";
import { logger } from "../utils/logger.ts";

const log = logger({ name: "Event: Ready" });

events.ready = () => {
  log.info("Bot Ready");
};`
}

func UtilsHelpersTsContent() string {
	return fmt.Sprintf(`import {
	Bot,
	BotWithCache,
	CreateApplicationCommand,
	getGuild,
	Guild,
	MakeRequired,
	upsertApplicationCommands,
} from "../../deps.ts";
import { logger } from "./logger.ts";
import { commands } from "../commands/mod.ts";
import { subCommand, subCommandGroup } from "../commands/mod.ts";

const log = logger({ name: "Helpers" });

/** This function will update all commands, or the defined scope */
export async function updateCommands(
	bot: BotWithCache,
	scope?: "Guild" | "Global"
) {
	const globalCommands: MakeRequired<CreateApplicationCommand, "name">[] = [];
	const perGuildCommands: MakeRequired<CreateApplicationCommand, "name">[] = [];

	for (const command of commands.values()) {
		if (command.scope) {
			if (command.scope === "Guild") {
				perGuildCommands.push({
					name: command.name,
					description: command.description,
					type: command.type,
					options: command.options ? command.options : undefined,
				});
			} else if (command.scope === "Global") {
					globalCommands.push({
					name: command.name,
					description: command.description,
					type: command.type,
					options: command.options ? command.options : undefined,
				});
			}
		} else {
			perGuildCommands.push({
				name: command.name,
				description: command.description,
				type: command.type,
				options: command.options ? command.options : undefined,
			});
		}
	}

	if (globalCommands.length && (scope === "Global" || scope === undefined)) {
		log.info(
			"Updating Global Commands, this takes up to 1 hour to take effect..."
		);

		await bot.helpers
			.upsertApplicationCommands(globalCommands)
			.catch(log.error);
	}

	if (perGuildCommands.length && (scope === "Guild" || scope === undefined)) {
		await bot.guilds.forEach(async (guild) => {
			await upsertApplicationCommands(bot, perGuildCommands, guild.id);
		});
	}
}

/** Update commands for a guild */
export async function updateGuildCommands(bot: Bot, guild: Guild) {
	const perGuildCommands: MakeRequired<CreateApplicationCommand, "name">[] = [];

	for (const command of commands.values()) {
		if (command.scope) {
			if (command.scope === "Guild") {
				perGuildCommands.push({
					name: command.name,
					description: command.description,
					type: command.type,
					options: command.options ? command.options : undefined,
				});
			}
		}
	}

	if (perGuildCommands.length) {
		await upsertApplicationCommands(bot, perGuildCommands, guild.id);
	}
}

export async function getGuildFromId(
	bot: BotWithCache,
	guildId: bigint
): Promise<Guild> {
	let returnValue: Guild = {} as Guild;

	if (guildId !== 0n) {
		if (bot.guilds.get(guildId)) {
			returnValue = bot.guilds.get(guildId) as Guild;
		}

		await getGuild(bot, guildId).then((guild) => {
			bot.guilds.set(guildId, guild);
			returnValue = guild;
		});
	}

	return returnValue;
}

export function snowflakeToTimestamp(id: bigint) {
	return Number(id / 4194304n + 1420070400000n);
}

export function humanizeMilliseconds(milliseconds: number) {
	// Gets ms into seconds
	const time = milliseconds / 1000;
	if (time < 1) return "1s";

	const days = Math.floor(time / 86400);
	const hours = Math.floor((%s) / 3600);
	const minutes = Math.floor(((%s) / 60);
	const seconds = Math.floor(((%s);

	const dayString = days ? %s : "";
	const hourString = hours ? %s : "";
	const minuteString = minutes ? %s : "";
	const secondString = seconds ? %s : "";

	return %s;
}

export function isSubCommand(
	data: subCommand | subCommandGroup
): data is subCommand {
	return !Reflect.has(data, "subCommands");
}

export function isSubCommandGroup(
	data: subCommand | subCommandGroup
): data is subCommandGroup {
	return Reflect.has(data, "subCommands");
}`, "time" + "%" + "86400", "time" + " % " + "86400" + ")" + " % " + "3600", "time" + " % " + "86400" + ")" + " % " + "3600" + ")" + " % " + "60" ,"`${days}d `", "`${hours}h `", "`${minutes}m `", "`${seconds}s `", "`${dayString}${hourString}${minuteString}${secondString}`", )
}

func UtilsLoggerTsContent() string {
	return fmt.Sprintf(`// deno-lint-ignore-file no-explicit-any
import { bold, cyan, gray, italic, red, yellow } from "../../deps.ts";

export enum LogLevels {
  Debug,
  Info,
  Warn,
  Error,
  Fatal,
}

const prefixes = new Map<LogLevels, string>([
  [LogLevels.Debug, "DEBUG"],
  [LogLevels.Info, "INFO"],
  [LogLevels.Warn, "WARN"],
  [LogLevels.Error, "ERROR"],
  [LogLevels.Fatal, "FATAL"],
]);

const noColor: (str: string) => string = (msg) => msg;
const colorFunctions = new Map<LogLevels, (str: string) => string>([
  [LogLevels.Debug, gray],
  [LogLevels.Info, cyan],
  [LogLevels.Warn, yellow],
  [LogLevels.Error, (str: string) => red(str)],
  [LogLevels.Fatal, (str: string) => red(bold(italic(str)))],
]);

export function logger({
  logLevel = LogLevels.Info,
  name,
}: {
  logLevel?: LogLevels;
  name?: string;
} = {}) {
  function log(level: LogLevels, ...args: any[]) {
    if (level < logLevel) return;

    let color = colorFunctions.get(level);
    if (!color) color = noColor;

    const date = new Date();
    const log = [
      %s,
      color(prefixes.get(level) || "DEBUG"),
      name ? %s : ">",
      ...args,
    ];

    switch (level) {
      case LogLevels.Debug:
        return console.debug(...log);
      case LogLevels.Info:
        return console.info(...log);
      case LogLevels.Warn:
        return console.warn(...log);
      case LogLevels.Error:
        return console.error(...log);
      case LogLevels.Fatal:
        return console.error(...log);
      default:
        return console.log(...log);
    }
  }

  function setLevel(level: LogLevels) {
    logLevel = level;
  }

  function debug(...args: any[]) {
    log(LogLevels.Debug, ...args);
  }

  function info(...args: any[]) {
    log(LogLevels.Info, ...args);
  }

  function warn(...args: any[]) {
    log(LogLevels.Warn, ...args);
  }

  function error(...args: any[]) {
    log(LogLevels.Error, ...args);
  }

  function fatal(...args: any[]) {
    log(LogLevels.Fatal, ...args);
  }

  return {
    log,
    setLevel,
    debug,
    info,
    warn,
    error,
    fatal,
  };
}

export const log = logger();`, "`[${date.toLocaleDateString()} ${date.toLocaleTimeString()}]`", "`${name} >`")
}

func Resources() string {
	return `# Botway Discord (Deno 🦕) Resources

> Here is some useful links and resources to help you to build your own bot

## Setup

- [Setup discord bot token](https://github.com/abdfnx/botway/discussions/4)
- [Get the guild id of your server](https://github.com/abdfnx/botway/discussions/4#discussioncomment-2653737)

## API

- [Discord API library for Deno](https://github.com/discordeno/discordeno)
- [Discordeno Website](https://discordeno.mod.land)
- [Discordeno Docs](https://doc.deno.land/https/deno.land/x/discordeno/mod.ts)
- [Discord Server](https://discord.com/invite/5vBgXk3UcZ)

## Examples

- [A collection of amazing examples written with Discordeno](https://github.com/discordeno/discordeno/tree/main/template)

big thanks to [**@discordeno**](https://github.com/discordeno) org`
}
