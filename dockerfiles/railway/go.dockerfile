FROM golang:alpine

COPY .botway.yaml .

ENV BOTWAY-DIR /root/.botway

RUN mkdir ${BOTWAY-DIR}

COPY botway.json ${BOTWAY-DIR}

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
