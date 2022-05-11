FROM alpine:latest
FROM python:alpine
FROM botwayorg/botway:latest

ENV PORT="8000"
ENV SLACK_BOT_NAME="{{.BotName}}"
ARG SLACK_TOKEN
ARG SLACK_APP_TOKEN

COPY . .

RUN apk update && \
	apk add --no-cache --virtual build-dependencies build-base gcc abuild binutils binutils-doc gcc-doc python3-dev libffi-dev git

RUN botway init --docker 
RUN curl https://raw.githubusercontent.com/pypa/pipenv/master/get-pipenv.py | python
RUN pipenv sync --system

# Add packages you want
# RUN apk add PACKAGE_NAME

EXPOSE $PORT

ENTRYPOINT [ "pipenv", "run", "python3", "-m" , "flask", "run", "--host=0.0.0.0", "--port", $PORT]
