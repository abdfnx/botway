FROM botwayorg/botway-cli:core AS core
FROM ubuntu:22.04

ENV BOTWAY_DIR /botway-dir/

RUN useradd --uid 1993 --user-group botway \
  && mkdir $BOTWAY_DIR \
  && chown botway:botway $BOTWAY_DIR

COPY --from=core /botway /usr/bin/botway
COPY ./docker/entry.sh /usr/local/bin/docker-entrypoint.sh

RUN chmod 755 /usr/local/bin/docker-entrypoint.sh

ENTRYPOINT ["docker-entrypoint.sh"]
CMD ["help"]
