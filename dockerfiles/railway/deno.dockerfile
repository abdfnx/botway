FROM denoland/deno:alpine

COPY . .

ENV BOTWAY-DIR /root/.botway

RUN mkdir ${BOTWAY-DIR}

COPY botway.json ${BOTWAY-DIR}

ENV PACKAGES "build-dependencies build-base gcc git ffmpeg curl binutils openssl-dev zlib-dev boost boost-dev"

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

# To add more packages
# RUN apk add PACKAGE_NAME

RUN deno cache deps.ts

ENTRYPOINT ["deno", "run", "--allow-all", "main.ts"]
