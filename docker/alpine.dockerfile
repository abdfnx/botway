FROM botwayorg/botway-cli:core AS core
FROM alpine:3.18.3

RUN apk update && apk add bash sudo

ENV BOTWAY_DIR /botway-dir/

RUN addgroup --gid 1000 botway \
  && adduser --uid 1000 --disabled-password botway --ingroup botway \
  && mkdir -p $BOTWAY_DIR \
  && chown botway:botway $BOTWAY_DIR

COPY --from=core /botway /bin/botway
COPY ./docker/entry.sh /usr/local/bin/docker-entrypoint.sh

RUN chmod 755 /usr/local/bin/docker-entrypoint.sh

ENTRYPOINT ["docker-entrypoint.sh"]
CMD [ "help" ]
