FROM botwayorg/botway:latest AS bw

ENV PACKAGES "build-dependencies build-base gcc git"

COPY . .

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

# Add packages you want
# RUN apk add PACKAGE_NAME

RUN botway init --docker

FROM golang:alpine

COPY --from=bw /root/.botway /root/.botway

WORKDIR /app/

COPY . .

RUN go mod tidy
RUN go build -o bot ./src/main.go

EXPOSE 8000

ENTRYPOINT ["./bot"]
