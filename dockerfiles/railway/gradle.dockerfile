FROM gradle:alpine

COPY . .

RUN mkdir /root/.botway

COPY botway.json /root/.botway

ENV PACKAGES "build-dependencies build-base openssl openssl-dev musl-dev libressl-dev gcc git libsodium opus ffmpeg m4 binutils zlib-dev boost boost-dev"

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

# To add more packages
# RUN apk add PACKAGE_NAME

RUN gradle wrapper

RUN gradle build --no-daemon

RUN ./gradlew

ENTRYPOINT [ "./gradlew", "run" ]
