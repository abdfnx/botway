FROM alpine:latest
FROM node:alpine
FROM denoland/deno:alpine
FROM botwayorg/botway:latest

ENV DISCORD_BOT_NAME="{{.BotName}}"
ARG DISCORD_TOKEN
ARG DISCORD_CLIENT_ID

ADD . .

RUN apk update && \
	apk add --no-cache --virtual build-dependencies build-base gcc git ffmpeg curl

# Add packages you want
# RUN apk add PACKAGE_NAME

RUN botway init --docker --name ${DISCORD_BOT_NAME}

USER deno

RUN deno cache deps.ts

EXPOSE 8000

ENTRYPOINT ["deno", "run", "--allow-net", "--allow-write", "--allow-read", "--allow-run", "mod.ts"]
