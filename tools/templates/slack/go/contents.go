package sgo

import (
	"github.com/abdfnx/botway/tools/templates"
)

func DockerfileContent(botName string) string {
	return templates.Content("slack", "go", "Dockerfile", botName)
}

func MainGoContent() string {
	return templates.Content("slack", "go", "src/main.go.bw", "")
}

func Resources() string {
	return `# Botway Slack (Go) Resources

> Here is some useful links and resources to help you to build your own bot

## Setup

- [Setup Slack bot](https://github.com/abdfnx/botway/discussions/5)

## API

- [Slack Bot Framework for Golang](https://github.com/shomali11/slacker)
- [Slack Golang API](https://github.com/slack-go/slack)

## Exmaples

- [Slacker examples](https://github.com/shomali11/slacker/tree/master/examples)

big thanks to [**@shomali11**](https://github.com/shomali11)`
}
