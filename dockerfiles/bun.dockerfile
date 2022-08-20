FROM botwayorg/botway:latest AS bw

COPY . .

RUN botway init --docker

FROM jarredsumner/bun:edge

COPY --from=bw /root/.botway /root/.botway

COPY . .

ENV PATH="/root/.bun/bin:$PATH"

RUN bun i

ENTRYPOINT [ "bun", "src/main.js" ]
