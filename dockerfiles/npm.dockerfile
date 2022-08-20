FROM botwayorg/botway:latest AS bw

COPY . .

RUN botway init --docker

FROM node:alpine

ENV NODE_ENV "production"
ENV PACKAGES "build-dependencies libtool autoconf automake gcc gcc-doc g++ make py3-pip py-pip zlib-dev python3 python3-dev libffi-dev build-base gcc git ffmpeg binutils openssl-dev zlib-dev boost boost-dev"

COPY --from=bw /root/.botway /root/.botway

COPY . .

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

# Add packages you want
# RUN apk add PACKAGE_NAME

RUN npm i --production

ENTRYPOINT ["node", "./src/main.js"]
