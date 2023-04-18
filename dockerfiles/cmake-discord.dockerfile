FROM botwayorg/botway-cli:latest AS bw

ARG {{.BotSecrets}}

COPY . .

RUN botway docker-init

FROM brainboxdotcc/dpp:latest

WORKDIR /usr/src/{{.BotName}}

COPY --from=bw /root/.botway /root/.botway

COPY . .

WORKDIR /usr/src/{{.BotName}}/build

RUN cmake ..

RUN make -j$(nproc)

ENTRYPOINT [ "/usr/src/{{.BotName}}/build/{{.BotName}}" ]
