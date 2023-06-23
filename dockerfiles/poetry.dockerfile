FROM botwayorg/botway-cli:latest AS bw

ARG {{.BotSecrets}}

COPY . .

RUN botway docker-init

FROM python:alpine

ENV PACKAGES "build-dependencies build-base gcc abuild binutils binutils-doc gcc-doc py-pip jpeg-dev zlib-dev python3 python3-dev libffi-dev git boost boost-dev"

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

# To add more packages
# RUN apk add PACKAGE_NAME

COPY --from=bw /root/.botway /root/.botway

COPY . .

# Install poetry
RUN pip install poetry

RUN poetry config virtualenvs.create false

RUN poetry install

ENTRYPOINT ["poetry", "run", "python3", "./src/main.py"]
