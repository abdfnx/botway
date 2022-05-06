package dgo

import "github.com/abdfnx/botway/templates"

func DockerfileContent(botName string) string {
	return templates.Content("discord", "go", "Dockerfile", botName)
}

func MainGoContent() string {
	return templates.Content("discord", "go", "src/main.go.bw", "")
}

func Resources() string {
	return `# Botway Discord (Go) Resources

> Here is some useful links and resources to help you to build your own bot

## Setup

- [Setup discord bot token](https://github.com/abdfnx/botway/discussions/4)
- [Get the guild id of your server](https://github.com/abdfnx/botway/discussions/4#discussioncomment-2653737)

## API

- [Discord bot api for golang](https://github.com/bwmarrin/discordgo)
- [Voice implementation for Discordgo](https://github.com/bwmarrin/dgvoice)
- [Specification & Tool the Discord Audio (dca) file format, supported by all the best Discord libs](https://github.com/bwmarrin/dca)
- [Discord Server](https://discord.gg/golang)

## Examples

- [A collection of example programs written with DiscordGo](https://github.com/bwmarrin/discordgo/tree/master/examples)
- [Examples for dgvoice](https://github.com/bwmarrin/dgvoice/tree/master/examples)

big thanks to [**@bwmarrin**](https://github.com/bwmarrin)`
}
