FROM crystallang/crystal:nightly-alpine

COPY . .

RUN mkdir /root/.botway

COPY botway.json /root/.botway

ENV PACKAGES "build-dependencies build-base gcc git libsodium opus ffmpeg binutils openssl-dev zlib-dev boost boost-dev"

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

# To add more packages
# RUN apk add PACKAGE_NAME

RUN shards install
RUN shards build --static --no-debug --release --production -v

ENTRYPOINT [ "./bin/{{.BotName}}" ]
