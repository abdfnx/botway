FROM dart:stable

COPY . .

RUN mkdir /root/.botway

COPY botway.json /root/.botway

RUN dart pub get

RUN dart compile exe src/main.dart -o bot

ENTRYPOINT ["./bot"]
