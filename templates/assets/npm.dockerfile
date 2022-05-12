FROM alpine:latest
FROM node:alpine
FROM botwayorg/botway:latest

ENV NODE_ENV "production"
ENV PACKAGES "build-dependencies build-base gcc git ffmpeg"

COPY . .

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

# Add packages you want
# RUN apk add PACKAGE_NAME

RUN botway init --docker
RUN npm i --production

EXPOSE 8000

ENTRYPOINT ["botway", "start"]
