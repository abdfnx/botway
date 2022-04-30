package tgo

import "fmt"

func DockerfileContent(botName string) string {
	return fmt.Sprintf(`FROM alpine:latest
FROM golang:alpine AS builder
FROM botwayorg/botway:latest

ENV DISCORD_BOT_NAME="%s"
ARG DISCORD_TOKEN
ARG DISCORD_CLIENT_ID

COPY . .

RUN apk update && \
	apk add --no-cache --virtual build-dependencies build-base gcc git

# Add packages you want
# RUN apk add PACKAGE_NAME

RUN botway init --docker --name ${DISCORD_BOT_NAME}
RUN go mod tidy
RUN go build src/main.go -o ${DISCORD_BOT_NAME}

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

EXPOSE 8000

ENTRYPOINT ["./${DISCORD_BOT_NAME}"]`, botName)
}

func MainGoContent() string {
	return `package main

import (
	"log"
	"time"

	"github.com/abdfnx/botwaygo"
	tele "gopkg.in/telebot.v3"
)

func main() {
	pref := tele.Settings{
		Token:  botwaygo.GetToken(),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	b.Start()
}`
}

func Resources() string {
	return `# Botway Telegtam (Go) Resources

> Here is some useful links and resources to help you to build your own bot

## Setup

- [Setup telegram bot](https://github.com/abdfnx/botway/discussions/5)

## API

- [A Telegram bot framework in Go.](https://github.com/tucnak/telebot)
- [Telegram Group](https://t.me/go_telebot)

big thanks to [**@python-telegram-bot**](https://github.com/python-telegram-bot) org`
}
