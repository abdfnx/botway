package nodejs

import "github.com/abdfnx/botway/tools/templates"

var Packages = "telegraf botway.js"

func IndexJSContent() string {
	return templates.Content("telegram", "nodejs", "src/index.js", "")
}

func Resources() string {
	return `# Botway Telegram (Node.js) Resources

> Here is some useful links and resources to help you to build your own bot

## Setup

- [Setup Telegram bot](https://github.com/abdfnx/botway/discussions/5)

## API

- [Modern Telegram Bot Framework for Node.js](https://github.com/telegraf/telegraf)
- [Telegraf Docs](https://github.com/telegraf/telegraf/tree/v4/docs)
- [Telegraf Telegram Channel](https://t.me/TelegrafJSChat)

## Examples

[Examples](https://github.com/telegraf/telegraf/tree/v4/docs/examples)

big thanks to [**@telegraf**](https://github.com/telegraf/telegraf) org`
}
