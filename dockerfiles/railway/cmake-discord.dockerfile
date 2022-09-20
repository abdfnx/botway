FROM brainboxdotcc/dpp:latest

COPY .botway.yaml .

RUN mkdir /root/.botway

COPY botway.json /root/.botway

WORKDIR /usr/src/{{.BotName}}

COPY . .

WORKDIR /usr/src/{{.BotName}}/build

RUN cmake ..
RUN make -j$(nproc)

ENTRYPOINT [ "/usr/src/{{.BotName}}/build/{{.BotName}}" ]
