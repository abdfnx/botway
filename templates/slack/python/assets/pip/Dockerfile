FROM alpine:latest
FROM python:alpine
FROM botwayorg/botway:latest

ENV PORT="8000"
ENV SLACK_BOT_NAME="{{.BotName}}"
ARG SLACK_TOKEN
ARG SLACK_APP_TOKEN

COPY . .

RUN apk update && \
	apk add --no-cache --virtual build-dependencies build-base gcc abuild binutils python3-dev libffi-dev git

RUN botway init --docker --name ${SLACK_BOT_NAME}
RUN pip3 install -r requirements.txt

# Add packages you want
# RUN apk add PACKAGE_NAME

EXPOSE $PORT

ENTRYPOINT [ "python3", "-m" , "flask", "run", "--host=0.0.0.0", "--port", $PORT]
