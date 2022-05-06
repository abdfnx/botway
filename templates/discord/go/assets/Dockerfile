FROM alpine:latest
FROM golang:alpine AS builder
FROM botwayorg/botway:latest

ENV DISCORD_BOT_NAME="{{.BotName}}"
ARG DISCORD_TOKEN
ARG DISCORD_CLIENT_ID

COPY . .

RUN apk update && \
	apk add --no-cache --virtual build-dependencies build-base gcc git

# Add packages you want
# RUN apk add PACKAGE_NAME

RUN botway init --docker --name ${DISCORD_BOT_NAME}
RUN go mod tidy
RUN go build src/main.go -o ${DISCORD_BOT_NAME}

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

EXPOSE 8000

ENTRYPOINT ["./${DISCORD_BOT_NAME}"]
