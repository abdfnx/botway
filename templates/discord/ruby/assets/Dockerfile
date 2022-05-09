FROM alpine:latest
FROM ruby:alpine
FROM botwayorg/botway:latest

ARG DISCORD_TOKEN
ARG DISCORD_CLIENT_ID
# Uncomment the comment below to add your guild id, like this syntax --> ARG MY_SERVER_GUILD_ID
# ARG SERVER_GUILD_ID

COPY . .

RUN apk update && \
	apk add --no-cache --virtual build-dependencies build-base gcc git libsodium ffmpeg

# Add packages you want
# RUN apk add PACKAGE_NAME

RUN botway init --docker
RUN gem update --system
RUN gem install bundler
RUN bundle install

EXPOSE 8000

ENTRYPOINT ["bundle", "exec", "ruby", "./src/main.rb"]
