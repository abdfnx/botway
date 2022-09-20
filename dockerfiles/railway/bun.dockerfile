FROM jarredsumner/bun:edge

COPY . .

ENV BOTWAY-DIR /root/.botway

RUN mkdir ${BOTWAY-DIR}

COPY botway.json ${BOTWAY-DIR}

ENV PATH="/root/.bun/bin:$PATH"

RUN bun i

ENTRYPOINT [ "bun", "src/main.js" ]
