FROM botwayorg/botway:latest AS bw

COPY . .

RUN botway init --docker

FROM denoland/deno:alpine

ENV PACKAGES "build-dependencies build-base gcc git ffmpeg curl binutils openssl-dev zlib-dev boost boost-dev"

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

# Add packages you want
# RUN apk add PACKAGE_NAME

COPY --from=bw /root/.botway /root/.botway

COPY . .

RUN deno cache deps.ts

ENTRYPOINT ["deno", "run", "--allow-all", "main.ts"]
