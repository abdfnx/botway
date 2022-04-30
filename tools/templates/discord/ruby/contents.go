package ruby

import "fmt"

func MainRbContent() string {
	return `# frozen_string_literal: true

require 'discordrb'
require 'botwayrb'

bot = Discordrb::Bot.new token: botwayrb.getToken()

bot.message(with_text: 'ping') do |event|
  event.respond 'pong!'
end

bot.run`
}

func DockerfileContent(botName string) string {
	return fmt.Sprintf(`FROM alpine:latest
FROM ruby:alpine
FROM botwayorg/botway:latest

ENV DISCORD_BOT_NAME="%s"
ARG DISCORD_TOKEN
ARG DISCORD_CLIENT_ID

COPY . .

RUN apk update && \
	apk add --no-cache --virtual build-dependencies build-base gcc git libsodium ffmpeg

# Add packages you want
# RUN apk add PACKAGE_NAME

RUN botway init --docker --name ${DISCORD_BOT_NAME}
RUN gem update --system
RUN gem install bundler
RUN bundle install

EXPOSE 8000

ENTRYPOINT ["bundle", "exec", "ruby", "./src/main.rb"]`, botName)
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
