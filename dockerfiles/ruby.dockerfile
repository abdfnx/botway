FROM botwayorg/botway-cli:latest AS bw

ARG {{.BotSecrets}}

COPY . .

RUN botway docker-init

FROM ruby:alpine

ENV PACKAGES "build-dependencies build-base gcc git libsodium opus ffmpeg binutils openssl-dev zlib-dev boost boost-dev"

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

# To add more packages
# RUN apk add PACKAGE_NAME

COPY --from=bw /root/.botway /root/.botway

COPY . .

RUN gem update --system

RUN gem install bundler

RUN bundle install

ENTRYPOINT ["bundle", "exec", "ruby", "./src/main.rb"]
