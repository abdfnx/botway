FROM botwayorg/botway-cli:latest AS bw

ARG {{.BotSecrets}}

COPY . .

RUN botway docker-init

FROM swift:latest

RUN apt-get update -y && \
    apt-get install -y libopus-dev opus-tools git gcc libffi-dev python-dev ffmpeg build-essential autoconf automake libtool m4 youtube-dl libcurl4-openssl-dev

# To add more packages
# RUN apt-get install -y PACKAGE_NAME

RUN swift build

ENTRYPOINT [ "./.build/debug/{{.BotName}}" ]
