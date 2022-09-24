FROM reo7sp/tgbot-cpp:latest

COPY .botway.yaml .

RUN mkdir /root/.botway

COPY botway.json /root/.botway

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

COPY . .

WORKDIR /usr/src/{{.BotName}}/build

RUN cmake ..
RUN make -j$(nproc)

ENTRYPOINT [ "/usr/src/{{.BotName}}/build/{{.BotName}}" ]
