package nodejs

import "github.com/abdfnx/botway/templates"

var Packages = "node-telegram-bot-api botway.js"

func IndexJSContent() string {
	return templates.Content("telegram", "nodejs", "src/index.js", "")
}

func BotGif() string {
	return templates.Content("telegram", "nodejs", "src/bot.gif", "")
}

func Resources() string {
	return `# Botway Telegram (Node.js) Resources

> Here is some useful links and resources to help you to build your own bot

## Setup

- [Setup Telegram bot](https://github.com/abdfnx/botway/discussions/5)

## API

- [Telegram Bot API for NodeJS](https://github.com/https://github.com/yagop/node-telegram-bot-api)
- [node-telegram-bot-api Docs](https://github.com/yagop/node-telegram-bot-api/tree/master/doc)
- [node-telegram-bot-api Help Information](https://github.com/yagop/node-telegram-bot-api/blob/master/doc/help.md)
- [Tutorials](https://github.com/yagop/node-telegram-bot-api/tree/master/doc/tutorials.md)
- [node-telegram-bot-api Telegram Channel](https://t.me/node_telegram_bot_api)
- [node-telegram-bot-api Telegram Group](https://t.me/ntbasupport)

## Examples

[Examples](https://github.com/yagop/node-telegram-bot-api/tree/master/examples)

big thanks to [**@yagop**](https://github.com/yagop)`
}
