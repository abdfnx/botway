FROM rust:alpine

ENV PACKAGES "build-dependencies build-base openssl openssl-dev musl-dev libressl-dev gcc git lld clang libsodium ffmpeg opus autoconf automake libtool"

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

RUN git clone https://github.com/abdfnx/botway
RUN cd installer

# Build linux target
RUN 