FROM botwayorg/botway-cli:latest AS bw

ARG {{.BotSecrets}}

COPY . .

RUN botway docker-init

FROM php:alpine
FROM composer

ENV PACKAGES "build-dependencies build-base openssl openssl-dev musl-dev libressl-dev gcc git libsodium ffmpeg opus autoconf automake binutils zlib-dev boost boost-dev"

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

# To add more packages
# RUN apk add PACKAGE_NAME

COPY --from=bw /root/.botway /root/.botway

COPY . .

RUN composer install

ENTRYPOINT [ "php", "src/main.php" ]
