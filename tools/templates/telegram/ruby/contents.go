package ruby

import "github.com/abdfnx/botway/tools/templates"

func MainRbContent() string {
	return templates.Content("telegram", "ruby", "src/main.rb", "")
}

func DockerfileContent(botName string) string {
	return templates.Content("telegram", "ruby", "Dockerfile", botName)
}

func Resources() string {
	return `# Botway Telegram (Ruby 💎) Resources

> Here is some useful links and resources to help you to build your own bot

## Setup

- [Setup telegram bot](https://github.com/abdfnx/botway/discussions/5)

## API

- [Ruby wrapper for Telegram's Bot API](https://github.com/atipugin/telegram-bot-ruby)

## Examples

- [Example](https://github.com/atipugin/telegram-bot-ruby/tree/master/examples)

big thanks to [**@atipugin**](https://github.com/atipugin)`
}
