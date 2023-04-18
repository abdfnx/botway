FROM botwayorg/botway-cli:latest AS bw

ARG {{.BotSecrets}}

COPY . .

RUN botway docker-init

FROM reo7sp/tgbot-cpp:latest

RUN apt-get update -y && \
    apt-get install -y libopus-dev opus-tools git gcc cmake make libffi-dev python-dev ffmpeg build-essential autoconf automake libtool m4 youtube-dl

# To add more packages
# RUN apt-get install -y PACKAGE_NAME

RUN git clone https://github.com/nlohmann/json && \
    cd json && \
    cmake . && \
    make -j && \
    make install

WORKDIR /usr/src/{{.BotName}}

COPY --from=bw /root/.botway /root/.botway

COPY . .

WORKDIR /usr/src/{{.BotName}}/build

RUN cmake ..

RUN make -j$(nproc)

ENTRYPOINT [ "/usr/src/{{.BotName}}/build/{{.BotName}}" ]
