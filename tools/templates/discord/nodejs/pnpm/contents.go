package pnpm

import "fmt"

func DockerfileContent(botName string) string {
	return fmt.Sprintf(`FROM alpine:latest
FROM node:alpine
FROM botwayorg/botway:latest

ENV DISCORD_BOT_NAME="%s"
ARG DISCORD_TOKEN
ARG DISCORD_CLIENT_ID

COPY . .

RUN apk update && \
	apk add --no-cache --virtual build-dependencies build-base gcc git ffmpeg

# Add packages you want
# RUN apk add PACKAGE_NAME

RUN botway init --docker --name ${DISCORD_BOT_NAME}
RUN curl -fsSL https://get.pnpm.io/install.sh | PNPM_VERSION=7.0.0-rc.9 sh -

RUN pnpm fetch --prod
ADD . .
RUN pnpm install -r --offline --prod

EXPOSE 8000

ENTRYPOINT ["node", "./src/index.js"]`, botName)
}
