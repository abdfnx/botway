package deno

import (
	"fmt"
	"os"
	"strings"

	"github.com/abdfnx/resto/core/api"
)

func Content(fileName, botName string) string {
	url := "https://raw.githubusercontent.com/abdfnx/botway/main/tools/templates/discord/deno/assets/" + fileName
	respone, status, _, err := api.BasicGet(url, "GET", "", "", "", "", false, 0, nil)

	if err != nil {
		fmt.Println(err.Error())
	}

	if status == "404" || status == "401" || strings.Contains(respone, "404") {
		fmt.Println("404")
		os.Exit(0)
	}

	if strings.Contains(fileName, "Dockerfile") {
		return strings.ReplaceAll(respone, "{{.Discord_Bot_name}}", botName)
	} else {
		return respone
	}
}

func DockerfileContent(botName string) string {
	return Content("Dockerfile", botName)
}

func ModTsContent() string {
	return Content("mod.ts", "")
}

func DepsTsContent() string {
	return Content("deps.ts", "")
}

func CommandsModTsContent() string {
	return Content("src/commands/mod.ts", "")
}

func CommandsPingTsContent() string {
	return Content("src/commands/ping.ts", "")
}

func EventsGuildCreateTsContent() string {
	return Content("src/events/guildCreate.ts", "")
}

func EventsInteractionCreateTsContent() string {
	return Content("src/events/interactionCreate.ts", "")
}

func EventsModTsContent() string {
	return Content("src/events/mod.ts", "")
}

func EventsReadyTsContent() string {
	return Content("src/events/ready.ts", "")
}

func UtilsHelpersTsContent() string {
	return Content("src/utils/helpers.ts", "")
}

func UtilsLoggerTsContent() string {
	return Content("src/utils/logger.ts", "")
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
