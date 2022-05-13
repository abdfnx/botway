FROM alpine:latest
FROM ruby:alpine
FROM botwayorg/botway:latest

ENV PACKAGES "build-dependencies build-base gcc git libsodium ffmpeg"

COPY . .

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

# Add packages you want
# RUN apk add PACKAGE_NAME

RUN botway init --docker
RUN gem update --system
RUN gem install bundler
RUN bundle install

EXPOSE 8000

ENTRYPOINT ["bundle", "exec", "ruby", "./src/main.rb"]
