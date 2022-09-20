FROM golang:alpine

COPY .botway.yaml .

RUN mkdir /root/.botway

COPY botway.json /root/.botway

ENV PACKAGES "build-dependencies build-base gcc git binutils openssl-dev zlib-dev boost boost-dev"

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

# To add more packages
# RUN apk add PACKAGE_NAME

WORKDIR /app/

COPY . .

RUN go mod tidy
RUN go build -o bot ./src/main.go

ENTRYPOINT ["./bot"]
