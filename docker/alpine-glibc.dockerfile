FROM botwayorg/botway:core AS core
FROM frolvlad/alpine-glibc:alpine-3.14_glibc-2.33

RUN apk update && apk add bash sudo curl

ENV BOTWAY_DIR /botway-dir/

RUN addgroup --gid 1000 botway \
  && adduser --uid 1000 --disabled-password botway --ingroup botway \
  && mkdir -p $BOTWAY_DIR \
  && chown botway:botway $BOTWAY_DIR

COPY --from=core /botway /bin/botway
COPY ./docker/entry.sh /usr/local/bin/docker-entrypoint.sh

RUN chmod 755 /usr/local/bin/docker-entrypoint.sh

### install deno ###
RUN curl -fsSL https://deno.land/install.sh | sh
ENV DENO_INSTALL="/root/.deno"
ENV PATH="${DENO_INSTALL}/bin:${PATH}"

ENTRYPOINT ["docker-entrypoint.sh"]
CMD [ "help" ]
