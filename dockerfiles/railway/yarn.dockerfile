FROM node:alpine

COPY . .

ENV BOTWAY-DIR /root/.botway

RUN mkdir ${BOTWAY-DIR}

COPY botway.json ${BOTWAY-DIR}

ENV PACKAGES "build-dependencies libtool autoconf automake gcc gcc-doc g++ make py3-pip py-pip zlib-dev python3 python3-dev libffi-dev build-base gcc git ffmpeg binutils openssl-dev zlib-dev boost boost-dev"

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

# To add more packages
# RUN apk add PACKAGE_NAME

RUN yarn

ENTRYPOINT ["yarn", "start"]
