FROM alpine:latest
FROM denoland/deno:alpine
FROM botwayorg/botway:latest

ARG DISCORD_TOKEN
ARG DISCORD_CLIENT_ID
# Uncomment the comment below to add your guild id, like this syntax --> ARG MY_SERVER_GUILD_ID
# ARG SERVER_GUILD_ID

ADD . .

RUN apk update && \
	apk add --no-cache --virtual build-dependencies build-base gcc git ffmpeg curl

# Add packages you want
# RUN apk add PACKAGE_NAME

RUN botway init --docker

USER deno

RUN deno cache deps.ts

EXPOSE 8000

ENTRYPOINT ["deno", "run", "--allow-all", "mod.ts"]
