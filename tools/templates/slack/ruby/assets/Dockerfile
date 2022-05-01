FROM alpine:latest
FROM ruby:alpine
FROM botwayorg/botway:latest

ENV SLACK_BOT_NAME="{{.BotName}}"
ARG SLACK_TOKEN
ARG SLACK_APP_TOKEN

COPY . .

RUN apk update && \
	apk add --no-cache --virtual build-dependencies build-base gcc git libsodium ffmpeg

# Add packages you want
# RUN apk add PACKAGE_NAME

RUN botway init --docker --name ${TELEGRAM_BOT_NAME}
RUN gem update --system
RUN gem install bundler
RUN bundle install

EXPOSE 8000

ENTRYPOINT ["bundle", "exec", "ruby", "./src/main.rb"]
