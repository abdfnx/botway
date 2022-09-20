FROM jarredsumner/bun:edge

COPY . .

RUN mkdir /root/.botway

COPY botway.json /root/.botway

ENV PATH="/root/.bun/bin:$PATH"

RUN bun i

ENTRYPOINT [ "bun", "src/main.js" ]
