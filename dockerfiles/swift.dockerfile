FROM botwayorg/botway:latest AS bw

COPY . .

RUN botway init --docker

FROM swift:latest

RUN apt-get update -y && \
    apt-get install -y libopus-dev opus-tools git gcc libffi-dev python-dev ffmpeg build-essential autoconf automake libtool m4 youtube-dl libcurl4-openssl-dev

RUN swift build

ENTRYPOINT [ "./.build/debug/{{.BotName}}" ]
