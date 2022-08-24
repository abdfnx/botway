FROM botwayorg/botway:latest AS bw

COPY . .

RUN botway init --docker

FROM golang:alpine

ENV PACKAGES "build-dependencies build-base gcc git binutils openssl-dev zlib-dev boost boost-dev"

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

# Add packages you want
# RUN apk add PACKAGE_NAME

COPY --from=bw /root/.botway /root/.botway

WORKDIR /app/

COPY . .

RUN go mod tidy
RUN go build -o bot ./src/main.go

ENTRYPOINT ["./bot"]
