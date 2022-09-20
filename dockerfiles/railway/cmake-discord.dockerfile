FROM brainboxdotcc/dpp:latest

COPY .botway.yaml .

ENV BOTWAY-DIR /root/.botway

RUN mkdir ${BOTWAY-DIR}

COPY botway.json ${BOTWAY-DIR}

WORKDIR /usr/src/{{.BotName}}

COPY . .

WORKDIR /usr/src/{{.BotName}}/build

RUN cmake ..
RUN make -j$(nproc)

ENTRYPOINT [ "/usr/src/{{.BotName}}/build/{{.BotName}}" ]
