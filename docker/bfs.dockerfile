FROM golang:alpine

RUN apk update && apk add git curl bash sudo

ENV BOTWAY_DIR /botway-dir/

RUN addgroup --gid 1000 botway \
  && adduser --uid 1000 --disabled-password botway --ingroup botway \
  && mkdir -p $BOTWAY_DIR \
  && chown botway:botway $BOTWAY_DIR

RUN go install github.com/go-task/task/v3/cmd/task@latest

WORKDIR /botway-src

COPY . .

RUN task bfs

COPY ./docker/entry.sh /usr/local/bin/docker-entrypoint.sh

RUN chmod 755 /usr/local/bin/docker-entrypoint.sh

ENTRYPOINT ["docker-entrypoint.sh"]
CMD [ "help" ]
