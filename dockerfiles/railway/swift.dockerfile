FROM swift:latest

COPY . .

ENV BOTWAY-DIR /root/.botway

RUN mkdir ${BOTWAY-DIR}

COPY botway.json ${BOTWAY-DIR}

RUN apt-get update -y && \
    apt-get install -y libopus-dev opus-tools git gcc libffi-dev python-dev ffmpeg build-essential autoconf automake libtool m4 youtube-dl libcurl4-openssl-dev

# To add more packages
# RUN apt-get install -y PACKAGE_NAME

RUN swift build

ENTRYPOINT [ "./.build/debug/{{.BotName}}" ]
