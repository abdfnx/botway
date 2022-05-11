FROM alpine:latest
FROM rust:alpine
FROM botwayorg/botway:latest

ENV TELEGRAM_BOT_NAME="{{.BotName}}"
ARG TELEGRAM_TOKEN

COPY . .

RUN apk update && \
	apk add --no-cache --virtual build-dependencies build-base gcc git curl libsodium ffmpeg opus autoconf automake libtool m4 youtube-dl

# Add packages you want
# RUN apk add PACKAGE_NAME

RUN botway init --docker 
RUN curl -L get.fleet.rs | sh
RUN fleet build --release
RUN cp ./target/release/${TELEGRAM_BOT_NAME} .

EXPOSE 8000

ENTRYPOINT ["./${TELEGRAM_BOT_NAME}"]
