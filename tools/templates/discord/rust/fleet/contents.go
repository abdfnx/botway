package fleet

import "fmt"

func DockerfileContent(botName string) string {
	return fmt.Sprintf(`FROM alpine:latest
FROM rust:alpine
FROM botwayorg/botway:latest

ENV DISCORD_BOT_NAME="%s"
ARG DISCORD_TOKEN
ARG DISCORD_CLIENT_ID

COPY . .

RUN apk update && \
	apk add --no-cache --virtual build-dependencies build-base gcc git curl libsodium ffmpeg opus autoconf automake libtool m4 youtube-dl

# Add packages you want
# RUN apk add PACKAGE_NAME

RUN botway init --docker --name ${DISCORD_BOT_NAME}
RUN curl -L get.fleet.rs | sh
RUN fleet build --release
RUN cp ./target/release/${DISCORD_BOT_NAME} .

EXPOSE 8000

ENTRYPOINT ["./${DISCORD_BOT_NAME}"]`, botName)
}
