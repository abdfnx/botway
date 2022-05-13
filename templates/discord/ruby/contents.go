package ruby

import "github.com/abdfnx/botway/templates"

func DockerfileContent(botName string) string {
	return templates.Content("assets/ruby.dockerfile", botName)
}

func MainRbContent() string {
	return templates.Content("discord/ruby/assets/main.rb", "")
}

func Resources() string {
	return `# Botway Discord (Ruby 💎) Resources

> Here is some useful links and resources to help you to build your own bot

## Setup

- [Setup discord bot token](https://github.com/abdfnx/botway/discussions/4)
- [Get the guild id of your server](https://github.com/abdfnx/botway/discussions/4#discussioncomment-2653737)

## API

- [Discord bot api for ruby](https://github.com/shardlab/discordrb)
- [Discord Server](https://discord.gg/cyK3Hjm)

## Examples

- [A collection of example programs written with Discordrb](https://github.com/shardlab/discordrb/tree/main/examples)

big thanks to [**@shardlab**](https://github.com/shardlab) org`
}
