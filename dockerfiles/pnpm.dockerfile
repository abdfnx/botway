FROM botwayorg/botway-cli:latest AS bw

ARG {{.BotSecrets}}

COPY . .

RUN botway docker-init

FROM node:alpine

ENV PACKAGES "build-dependencies libtool autoconf automake gcc gcc-doc g++ make py3-pip py-pip zlib-dev python3 python3-dev libffi-dev build-base gcc git ffmpeg binutils openssl-dev zlib-dev boost boost-dev"

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

# To add more packages
# RUN apk add PACKAGE_NAME

COPY --from=bw /root/.botway /root/.botway

COPY . .

RUN npm i -g pnpm

RUN pnpm fetch --prod

RUN pnpm install

ENTRYPOINT ["pnpm", "start"]
