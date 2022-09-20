FROM php:alpine
FROM composer

COPY . .

RUN mkdir /root/.botway

COPY botway.json /root/.botway

ENV PACKAGES "build-dependencies build-base openssl openssl-dev musl-dev libressl-dev gcc git libsodium ffmpeg opus autoconf automake binutils zlib-dev boost boost-dev"

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

# To add more packages
# RUN apk add PACKAGE_NAME

RUN composer install

ENTRYPOINT [ "php", "src/main.php" ]
