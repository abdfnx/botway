FROM swift:latest

COPY . .

RUN mkdir /root/.botway

COPY botway.json /root/.botway

RUN apt-get update -y && \
    apt-get install -y libopus-dev opus-tools git gcc libffi-dev python-dev ffmpeg build-essential autoconf automake libtool m4 youtube-dl libcurl4-openssl-dev

# To add more packages
# RUN apt-get install -y PACKAGE_NAME

RUN swift build

ENTRYPOINT [ "./.build/debug/{{.BotName}}" ]
