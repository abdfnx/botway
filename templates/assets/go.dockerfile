FROM alpine:latest
FROM golang:alpine AS builder
FROM botwayorg/botway:latest

ENV BOT_NAME "{{.BotName}}"
ENV PACKAGES "build-dependencies build-base gcc git"

COPY . .

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

# Add packages you want
# RUN apk add PACKAGE_NAME

RUN botway init --docker
RUN go mod tidy
RUN go build src/main.go -o bot

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

EXPOSE 8000

ENTRYPOINT ["./${BOT_NAME}"]
