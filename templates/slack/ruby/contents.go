package ruby

import "github.com/abdfnx/botway/templates"

func MainRbContent() string {
	return templates.Content("slack/ruby/assets/main.rb", "")
}

func DockerfileContent(botName string) string {
	return templates.Content("assets/ruby.dockerfile", botName)
}

func Resources() string {
	return `# Botway Slack (Ruby 💎) Resources

> Here is some useful links and resources to help you to build your own bot

## Setup

- [Setup Slack bot](https://github.com/abdfnx/botway/discussions/31)

## API

- [The easiest way to write a Slack bot in Ruby](https://github.com/slack-ruby/slack-ruby-bot)
- [A Ruby and command-line client for the Slack Web, Real Time Messaging and Event APIs.](https://github.com/slack-ruby/slack-ruby-client)

## Examples

- [Examples](https://github.com/slack-ruby/slack-ruby-bot/tree/master/examples)

big thanks to [**@slack-ruby**](https://github.com/slack-ruby) org`
}
