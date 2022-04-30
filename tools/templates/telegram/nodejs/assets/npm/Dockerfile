FROM alpine:latest
FROM node:alpine
FROM botwayorg/botway:latest

ENV NODE_ENV=production
ENV TELEGRAM_BOT_NAME="{{.BotName}}"
ARG TELEGRAM_TOKEN

COPY . .

RUN apk update && \
	apk add --no-cache --virtual build-dependencies build-base gcc git ffmpeg

# Add packages you want
# RUN apk add PACKAGE_NAME

RUN botway init --docker --name ${TELEGRAM_BOT_NAME}
RUN npm i --production

EXPOSE 8000

ENTRYPOINT ["node", "./src/index.js"]
