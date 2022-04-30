FROM alpine:latest
FROM python:alpine
FROM botwayorg/botway:latest

ENV DISCORD_BOT_NAME="{{.BotName}}"
ARG DISCORD_TOKEN
ARG DISCORD_CLIENT_ID

COPY . .

RUN apk update && \
	apk add --no-cache --virtual build-dependencies build-base gcc abuild binutils binutils-doc gcc-doc python3-dev libffi-dev git

RUN botway init --docker --name ${DISCORD_BOT_NAME}
RUN pip3 install -r requirements.txt

# Add packages you want
# RUN apk add PACKAGE_NAME

EXPOSE 8000

ENTRYPOINT ["python3", "./src/main.py"]
