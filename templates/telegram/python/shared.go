package python

import "github.com/abdfnx/botway/templates"

func MainPyContent() string {
	return templates.Content("telegram/python/assets/src/main.py", "")
}

func Resources() string {
	return `# Botway Telegtam (Python 🐍) Resources

> Here is some useful links and resources to help you to build your own bot

## Setup

- [Setup telegram bot](https://github.com/abdfnx/botway/discussions/5)

## API

- [Telegram Bot API for Python](https://github.com/python-telegram-bot/python-telegram-bot)
- [Python-Telegram-Bot Website](https://python-telegram-bot.org)
- [Telegram Group](https://telegram.me/pythontelegrambotgroup)

## Examples

- [A collection of examples written with Python-Telegram-Bot](https://github.com/python-telegram-bot/python-telegram-bot/tree/master/examples)

big thanks to [**@python-telegram-bot**](https://github.com/python-telegram-bot) org`
}
