FROM botwayorg/botway:latest

ENV PACKAGES "build-dependencies libtool autoconf automake gcc gcc-doc python2 g++ make py3-pip py-pip zlib-dev python3 python3-dev libffi-dev build-base npm gcc git ffmpeg"

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

# Add packages you want
# RUN apk add PACKAGE_NAME

ADD . .

RUN botway init --docker
RUN npm i -g pnpm@latest node-gyp

RUN /bin/bash -c "bash"

RUN pnpm fetch --prod
RUN pnpm install

EXPOSE 8000

ENTRYPOINT [ "node", "./src/index.js" ]
