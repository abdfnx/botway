package deno

import "github.com/abdfnx/botway/templates"

func DockerfileContent(botName string) string {
	return templates.Content("assets/deno.dockerfile", botName)
}

func ModTsContent() string {
	return templates.Content("discord/deno/assets/mod.ts", "")
}

func DepsTsContent() string {
	return templates.Content("discord/deno/assets/deps.ts", "")
}

func CommandsModTsContent() string {
	return templates.Content("discord/deno/assets/src/commands/mod.ts", "")
}

func CommandsPingTsContent() string {
	return templates.Content("discord/deno/assets/src/commands/ping.ts", "")
}

func EventsGuildCreateTsContent() string {
	return templates.Content("discord/deno/assets/src/events/guildCreate.ts", "")
}

func EventsInteractionCreateTsContent() string {
	return templates.Content("discord/deno/assets/src/events/interactionCreate.ts", "")
}

func EventsModTsContent() string {
	return templates.Content("discord/deno/assets/src/events/mod.ts", "")
}

func EventsReadyTsContent() string {
	return templates.Content("discord/deno/assets/src/events/ready.ts", "")
}

func UtilsHelpersTsContent() string {
	return templates.Content("discord/deno/assets/src/utils/helpers.ts", "")
}

func UtilsLoggerTsContent() string {
	return templates.Content("discord/deno/assets/src/utils/logger.ts", "")
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
