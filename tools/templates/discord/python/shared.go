package python

import (
	"fmt"
	"os"
	"strings"

	"github.com/abdfnx/resto/core/api"
)

func Content(fileName, botName string) string {
	url := "https://raw.githubusercontent.com/abdfnx/botway/main/tools/templates/discord/python/assets/" + fileName
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

func MainPyContent() string {
	return Content("src/main.py", "")
}

func Resources() string {
	return `# Botway Discord (Python 🐍) Resources

> Here is some useful links and resources to help you to build your own bot

## Setup

- [Setup discord bot token](https://github.com/abdfnx/botway/discussions/4)
- [Get the guild id of your server](https://github.com/abdfnx/botway/discussions/4#discussioncomment-2653737)

## API

- [An API wrapper for Discord written in Python.](https://github.com/Rapptz/discord.py)
- [Discord.py Website](https://discordpy.rtfd.org/en/latest)
- [Discord Server](https://discord.gg/r3sSKJJ)

## Examples

- [A collection of example programs written with Discord.py](https://github.com/Rapptz/discord.py/tree/master/examples)

big thanks to [**@Rapptz**](https://github.com/Rapptz) org`
}
