package deno

import "github.com/abdfnx/botway/templates"

func DockerfileContent(botName string) string {
	return templates.Content("deno.dockerfile", "dockerfiles", botName)
}

func Resources() string {
	return templates.Content("discord/deno.md", "resources", "")
}

func MainTsContent() string {
	return templates.Content("main.ts", "discord-deno", "")
}

func DepsTsContent() string {
	return templates.Content("deps.ts", "discord-deno", "")
}

func CommandsModTsContent() string {
	return templates.Content("src/commands/mod.ts", "discord-deno", "")
}

func CommandsPingTsContent() string {
	return templates.Content("src/commands/ping.ts", "discord-deno", "")
}

func EventsGuildCreateTsContent() string {
	return templates.Content("src/events/guildCreate.ts", "discord-deno", "")
}

func EventsInteractionCreateTsContent() string {
	return templates.Content("src/events/interactionCreate.ts", "discord-deno", "")
}

func EventsModTsContent() string {
	return templates.Content("src/events/mod.ts", "discord-deno", "")
}

func EventsReadyTsContent() string {
	return templates.Content("src/events/ready.ts", "discord-deno", "")
}

func UtilsHelpersTsContent() string {
	return templates.Content("src/utils/helpers.ts", "discord-deno", "")
}

func UtilsLoggerTsContent() string {
	return templates.Content("src/utils/logger.ts", "discord-deno", "")
}
