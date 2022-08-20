FROM botwayorg/botway:latest AS bw

COPY . .

RUN botway init --docker

FROM gradle:alpine

ENV PACKAGES "build-dependencies build-base openssl openssl-dev musl-dev libressl-dev gcc git libsodium opus ffmpeg m4 binutils zlib-dev boost boost-dev"

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

# Add packages you want
# RUN apk add PACKAGE_NAME

COPY --from=bw /root/.botway /root/.botway

COPY . .

RUN gradle wrapper

RUN gradle build --no-daemon

RUN ./gradlew

ENTRYPOINT [ "./gradlew", "run" ]
