FROM botwayorg/botway:latest

ENV PACKAGES "build-dependencies libtool autoconf automake gcc gcc-doc python2 g++ make py3-pip py-pip zlib-dev python3 python3-dev libffi-dev build-base npm gcc git ffmpeg"

COPY . .

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

# Add packages you want
# RUN apk add PACKAGE_NAME

RUN botway init --docker
RUN yarn

EXPOSE 8000

ENTRYPOINT ["node", "./src/main.js"]
