FROM alpine:latest
FROM golang:alpine AS builder
FROM botwayorg/botway:latest

ENV TELEGRAM_BOT_NAME="{{.BotName}}"
ARG TELEGRAM_TOKEN

COPY . .

RUN apk update && \
	apk add --no-cache --virtual build-dependencies build-base gcc git

# Add packages you want
# RUN apk add PACKAGE_NAME

RUN botway init --docker 
RUN go mod tidy
RUN go build src/main.go -o ${TELEGRAM_BOT_NAME}

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

EXPOSE 8000

ENTRYPOINT ["./${TELEGRAM_BOT_NAME}"]
