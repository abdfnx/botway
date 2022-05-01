package nodejs

import "github.com/abdfnx/botway/tools/templates"

var Packages = "slackbots botway.js"

func IndexJSContent() string {
	return templates.Content("slack", "nodejs", "src/index.js", "")
}

func Resources() string {
	return `# Botway Slack (Node.js) Resources

> Here is some useful links and resources to help you to build your own bot

## Setup

- [Setup Slack bot](https://github.com/abdfnx/botway/discussions/6)

## API

- [Simple way to control your Slack Bot](https://github.com/mishk0/slack-bot-api)

## Examples

[Examples](https://github.com/mishk0/slack-bot-api/tree/master/test)

big thanks to [**@mishk0**](https://github.com/mishk0)`
}
