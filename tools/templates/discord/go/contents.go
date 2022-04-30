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
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/abdfnx/botwaygo"
	"github.com/bwmarrin/discordgo"
)

func main() {
	dg, err := discordgo.New("Bot " + botwaygo.GetToken())

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()

	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.\nPress CTRL-C to exit.")

	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// In this example, we only care about messages that are "ping".
	if m.Content != "ping" {
		return
	}

	// We create the private channel with the user who sent the message.
	channel, err := s.UserChannelCreate(m.Author.ID)

	if err != nil {
		// If an error occurred, we failed to create the channel.
		//
		// Some common causes are:
		// 1. We don't share a server with the user (not possible here).
		// 2. We opened enough DM channels quickly enough for Discord to
		//    label us as abusing the endpoint, blocking us from opening
		//    new ones.
		fmt.Println("error creating channel:", err)

		s.ChannelMessageSend(
			m.ChannelID,
			"Something went wrong while sending the DM!",
		)

		return
	}

	// Then we send the message through the channel we created.
	_, err = s.ChannelMessageSend(channel.ID, "Pong!")

	if err != nil {
		// If an error occurred, we failed to send the message.
		//
		// It may occur either when we do not share a server with the
		// user (highly unlikely as we just received a message) or
		// the user disabled DM in their settings (more likely).
		fmt.Println("error sending DM message:", err)
		s.ChannelMessageSend(
			m.ChannelID,
			"Failed to send you a DM. "+
				"Did you disable DM in your privacy settings?",
		)
	}
}`
}

func Resources() string {
	return `# Botway Discord (Go) Resources

> Here is some useful links and resources to help you to build your own bot

## Setup

- [Setup discord bot token](https://github.com/abdfnx/botway/discussions/4)
- [Get the guild id of your server](https://github.com/abdfnx/botway/discussions/4#discussioncomment-2653737)

## API

- [Discord bot api for golang](https://github.com/bwmarrin/discordgo)
- [Voice implementation for Discordgo](https://github.com/bwmarrin/dgvoice)
- [Specification & Tool the Discord Audio (dca) file format, supported by all the best Discord libs](https://github.com/bwmarrin/dca)
- [Discord Server](https://discord.gg/golang)

## Examples

- [A collection of example programs written with DiscordGo](https://github.com/bwmarrin/discordgo/tree/master/examples)
- [Examples for dgvoice](https://github.com/bwmarrin/dgvoice/tree/master/examples)

big thanks to [**@bwmarrin**](https://github.com/bwmarrin)`
}
