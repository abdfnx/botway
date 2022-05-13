FROM alpine:latest
FROM denoland/deno:alpine
FROM botwayorg/botway:latest

ADD . .

RUN apk update && \
	apk add --no-cache --virtual build-dependencies build-base gcc git ffmpeg curl

# Add packages you want
# RUN apk add PACKAGE_NAME

RUN botway init --docker

USER deno

RUN deno cache deps.ts

EXPOSE 8000

ENTRYPOINT ["deno", "run", "--allow-all", "./src/main.ts"]
