package tgo

import (
	"github.com/abdfnx/botway/templates"
)

func DockerfileContent(botName string) string {
	return templates.Content("assets/go.dockerfile", botName)
}

func MainGoContent() string {
	return templates.Content("telegram/go/assets/main.go.bw", "")
}

func Resources() string {
	return `# Botway Telegtam (Go) Resources

> Here is some useful links and resources to help you to build your own bot

## Setup

- [Setup telegram bot](https://github.com/abdfnx/botway/discussions/5)

## API

- [A Telegram bot framework in Go](https://github.com/tucnak/telebot)
- [Telegram Group](https://t.me/go_telebot)

big thanks to [**@python-telegram-bot**](https://github.com/python-telegram-bot) org`
}
