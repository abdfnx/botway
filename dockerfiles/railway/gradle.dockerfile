FROM gradle:alpine

COPY . .

ENV BOTWAY-DIR /root/.botway

RUN mkdir ${BOTWAY-DIR}

COPY botway.json ${BOTWAY-DIR}

ENV PACKAGES "build-dependencies build-base openssl openssl-dev musl-dev libressl-dev gcc git libsodium opus ffmpeg m4 binutils zlib-dev boost boost-dev"

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

# To add more packages
# RUN apk add PACKAGE_NAME

RUN gradle wrapper

RUN gradle build --no-daemon

RUN ./gradlew

ENTRYPOINT [ "./gradlew", "run" ]
