FROM rust:alpine
FROM botwayorg/botway:latest

ENV BOT_NAME "{{.BotName}}"
ENV PACKAGES "build-dependencies build-base gcc git libsodium ffmpeg opus autoconf automake libtool m4 youtube-dl"

COPY . .

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

# Add packages you want
# RUN apk add PACKAGE_NAME

RUN botway init --docker
RUN cargo build --release
RUN cp ./target/release/${BOT_NAME} .

EXPOSE 8000

ENTRYPOINT ["./${BOT_NAME}"]
