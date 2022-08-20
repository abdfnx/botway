FROM botwayorg/botway:latest AS bw

COPY . .

RUN botway init --docker

FROM ruby:alpine

ENV PACKAGES "build-dependencies build-base gcc git libsodium opus ffmpeg binutils openssl-dev zlib-dev boost boost-dev"

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

# Add packages you want
# RUN apk add PACKAGE_NAME

COPY --from=bw /root/.botway /root/.botway

COPY . .

RUN gem update --system
RUN gem install bundler
RUN bundle install

ENTRYPOINT ["bundle", "exec", "ruby", "./src/main.rb"]
