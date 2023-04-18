FROM botwayorg/botway-cli:latest AS bw

ARG {{.BotSecrets}}

COPY . .

RUN botway docker-init

FROM python:alpine

ENV PACKAGES "build-dependencies build-base gcc abuild binutils py-pip binutils-doc gcc-doc python3-dev libffi-dev git binutils openssl-dev zlib-dev boost boost-dev"

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

# To add more packages
# RUN apk add PACKAGE_NAME

COPY --from=bw /root/.botway /root/.botway

COPY . .

RUN pip3 install -r requirements.txt

ENTRYPOINT ["python3", "./src/main.py"]
