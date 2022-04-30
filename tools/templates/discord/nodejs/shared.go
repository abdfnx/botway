package nodejs

import (
	"fmt"
	"os"
	"strings"

	"github.com/abdfnx/resto/core/api"
)

var Packages = "discord.js @discordjs/rest @discordjs/builders discord-api-types discord-rpc zlib-sync erlpack bufferutil utf-8-validate @discordjs/voice libsodium-wrappers @discordjs/opus sodium botway.js"

func Content(fileName, botName string) string {
	url := "https://raw.githubusercontent.com/abdfnx/botway/main/tools/templates/discord/nodejs/assets/" + fileName
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

func IndexJSContent() string {
	return Content("src/index.js", "")
}

func Resources() string {
	return `# Botway Discord (Node.js) Resources

> Here is some useful links and resources to help you to build your own bot

## Setup

- [Setup discord bot token](https://github.com/abdfnx/botway/discussions/4)
- [Get the guild id of your server](https://github.com/abdfnx/botway/discussions/4#discussioncomment-2653737)

## API

- [A powerful JavaScript library for interacting with the Discord API](https://github.com/discordjs/discord.js)
- [Discord.js Website](https://discord.js.org)
- [Discord.js Docs](https://discord.js.org/#/docs)
- [Discord.js Guide](https://discordjs.guide)
- [Discord Server](https://discord.gg/djs)

big thanks to [**@discordjs**](https://github.com/discordjs) org`
}
