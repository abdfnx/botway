FROM botwayorg/botway-cli:core AS core
FROM centos:8

ENV BOTWAY_DIR /botway-dir/

RUN groupadd -g 1993 botway \
  && adduser -u 1993 -g botway botway \
  && mkdir $BOTWAY_DIR \
  && chown botway:botway $BOTWAY_DIR

COPY --from=core /botway /bin/botway

COPY ./docker/entry.sh /usr/local/bin/docker-entrypoint.sh

RUN chmod 755 /usr/local/bin/docker-entrypoint.sh

ENTRYPOINT ["docker-entrypoint.sh"]
CMD ["help"]
