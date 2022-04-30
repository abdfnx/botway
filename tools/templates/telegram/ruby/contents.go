package ruby

import "fmt"

func MainRbContent() string {
	return `# frozen_string_literal: true

require 'telegram-bot-ruby'
require 'botwayrb'

Telegram::Bot::Client.run(botwayrb.getToken()) do |bot|
  bot.listen do |message|
    case message.text
    when '/start'
      bot.api.send_message(chat_id: message.chat.id, text: "Hello, #{message.from.first_name}")
    when '/stop'
      bot.api.send_message(chat_id: message.chat.id, text: "Bye, #{message.from.first_name}")
    end
  end
end`
}

func DockerfileContent(botName string) string {
	return fmt.Sprintf(`FROM alpine:latest
FROM ruby:alpine
FROM botwayorg/botway:latest

ENV TELEGRAM_BOT_NAME="%s"
ARG TELEGRAM_TOKEN

COPY . .

RUN apk update && \
	apk add --no-cache --virtual build-dependencies build-base gcc git libsodium ffmpeg

# Add packages you want
# RUN apk add PACKAGE_NAME

RUN botway init --docker --name ${TELEGRAM_BOT_NAME}
RUN gem update --system
RUN gem install bundler
RUN bundle install

EXPOSE 8000

ENTRYPOINT ["bundle", "exec", "ruby", "./src/main.rb"]`, botName)
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
