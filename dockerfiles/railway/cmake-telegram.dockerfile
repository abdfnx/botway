FROM reo7sp/tgbot-cpp:latest

COPY .botway.yaml .

ENV BOTWAY-DIR /root/.botway

RUN mkdir ${BOTWAY-DIR}

COPY botway.json ${BOTWAY-DIR}

RUN apt-get update -y && \
    apt-get install -y libopus-dev opus-tools git gcc cmake make libffi-dev python-dev ffmpeg build-essential autoconf automake libtool m4 youtube-dl

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
