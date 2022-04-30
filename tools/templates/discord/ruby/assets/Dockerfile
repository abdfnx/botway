FROM alpine:latest
FROM ruby:alpine
FROM botwayorg/botway:latest

ENV DISCORD_BOT_NAME="{{.Discord_Bot_name}}"
ARG DISCORD_TOKEN
ARG DISCORD_CLIENT_ID

COPY . .

RUN apk update && \
	apk add --no-cache --virtual build-dependencies build-base gcc git libsodium ffmpeg

# Add packages you want
# RUN apk add PACKAGE_NAME

RUN botway init --docker --name ${DISCORD_BOT_NAME}
RUN gem update --system
RUN gem install bundler
RUN bundle install

EXPOSE 8000

ENTRYPOINT ["bundle", "exec", "ruby", "./src/main.rb"]
