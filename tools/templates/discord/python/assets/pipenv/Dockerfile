FROM alpine:latest
FROM python:alpine
FROM botwayorg/botway:latest

ENV DISCORD_BOT_NAME="{{.Discord_Bot_name}}"
ARG DISCORD_TOKEN
ARG DISCORD_CLIENT_ID

COPY . .

RUN apk update && \
	apk add --no-cache --virtual build-dependencies build-base gcc abuild binutils binutils-doc gcc-doc python3-dev libffi-dev git

# Add packages you want
# RUN apk add PACKAGE_NAME

RUN botway init --docker --name ${DISCORD_BOT_NAME}
RUN curl https://raw.githubusercontent.com/pypa/pipenv/master/get-pipenv.py | python
RUN pipenv sync --system

EXPOSE 8000

ENTRYPOINT ["pipenv", "run", "python3", "./src/main.py"]
