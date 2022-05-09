FROM alpine:latest
FROM node:alpine
FROM botwayorg/botway:latest

ARG DISCORD_TOKEN
ARG DISCORD_CLIENT_ID
# Uncomment the comment below to add your guild id, like this syntax --> ARG MY_SERVER_GUILD_ID
# ARG SERVER_GUILD_ID

RUN apk update && \
	apk add --no-cache --virtual build-dependencies build-base gcc git ffmpeg

# Add packages you want
# RUN apk add PACKAGE_NAME

RUN botway init --docker
RUN curl -fsSL https://get.pnpm.io/install.sh | PNPM_VERSION=7.0.0 sh -

RUN pnpm fetch --prod
ADD . .
RUN pnpm install -r --offline --prod

EXPOSE 8000

ENTRYPOINT ["node", "./src/index.js"]
