FROM alpine:latest
FROM node:alpine
FROM botwayorg/botway:latest

ENV PACKAGES "build-dependencies build-base gcc git ffmpeg"
ENV PNPM_VERSION "7.0.0"

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

# Add packages you want
# RUN apk add PACKAGE_NAME

RUN botway init --docker
RUN curl -fsSL https://get.pnpm.io/install.sh | sh -

RUN pnpm fetch --prod
ADD . .
RUN pnpm install -r --offline --prod

EXPOSE 8000

ENTRYPOINT ["botway", "start"]
