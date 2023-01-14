FROM botwayorg/botway:latest AS bw

ARG {{.BotSecrets}}

COPY . .

RUN botway docker-init

FROM botwayorg/concord

# To add more packages
# RUN apk add PACKAGE_NAME

COPY --from=bw /root/.botway /root/.botway

COPY . .

RUN gcc src/main.c -o bot -pthread -ldiscord -lcurl

ENTRYPOINT [ "./bot" ]
