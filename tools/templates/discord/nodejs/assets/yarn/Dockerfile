FROM alpine:latest
FROM node:alpine
FROM botwayorg/botway:latest

ENV DISCORD_BOT_NAME="{{.Discord_Bot_name}}"
ARG DISCORD_TOKEN
ARG DISCORD_CLIENT_ID

COPY . .

RUN apk update && \
	apk add --no-cache --virtual build-dependencies build-base gcc git ffmpeg

# Add packages you want
# RUN apk add PACKAGE_NAME

RUN botway init --docker --name ${DISCORD_BOT_NAME}
RUN yarn

EXPOSE 8000

ENTRYPOINT ["node", "./src/index.js"]
