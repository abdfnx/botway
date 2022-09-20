FROM dart:stable

COPY . .

ENV BOTWAY-DIR /root/.botway

RUN mkdir ${BOTWAY-DIR}

COPY botway.json ${BOTWAY-DIR}

RUN dart pub get

RUN dart compile exe src/main.dart -o bot

ENTRYPOINT ["./bot"]
