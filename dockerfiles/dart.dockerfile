FROM botwayorg/botway:latest AS bw

COPY . .

RUN botway init --docker

FROM dart:stable

COPY --from=bw /root/.botway /root/.botway

COPY . .

RUN dart pub get

RUN dart compile exe src/main.dart -o bot

ENTRYPOINT ["./bot"]
