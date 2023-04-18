FROM botwayorg/botway-cli:latest AS bw

ARG {{.BotSecrets}}

COPY . .

RUN botway docker-init

FROM dart:stable

COPY --from=bw /root/.botway /root/.botway

COPY . .

RUN dart pub get

RUN dart compile exe src/main.dart -o bot

ENTRYPOINT ["./bot"]
