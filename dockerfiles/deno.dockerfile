FROM botwayorg/botway:alpine-glibc

ENV PACKAGES "build-dependencies build-base gcc git ffmpeg curl binutils openssl-dev zlib-dev boost boost-dev"

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

# Add packages you want
# RUN apk add PACKAGE_NAME

COPY . .

RUN botway init --docker

RUN deno cache deps.ts

ENTRYPOINT ["deno", "run", "--allow-all", "main.ts"]
