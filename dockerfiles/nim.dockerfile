FROM botwayorg/botway:latest AS bw

COPY . .

RUN botway init --docker

FROM nimlang/nim:alpine

ENV PACKAGES "build-dependencies build-base openssl openssl-dev musl-dev libressl-dev gcc git lld clang libsodium ffmpeg opus autoconf automake libtool m4 youtube-dl binutils zlib-dev boost boost-dev"

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

# Add packages you want
# RUN apk add PACKAGE_NAME

COPY . .

RUN nimble install -y

RUN nim c -d:ssl ./src/main.nim

ENTRYPOINT [ "./src/main" ]
