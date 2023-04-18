FROM botwayorg/botway-cli:latest AS bw

ARG {{.BotSecrets}}

COPY . .

RUN botway docker-init

FROM jarredsumner/bun:edge

COPY --from=bw /root/.botway /root/.botway

COPY . .

ENV PATH="/root/.bun/bin:$PATH"

RUN bun i

ENTRYPOINT [ "bun", "src/main.js" ]
