FROM botwayorg/botway:latest AS bw

COPY . .

RUN botway init --docker

FROM botwayorg/concord

# Add packages you want
# RUN apk add PACKAGE_NAME

COPY --from=bw /root/.botway /root/.botway

COPY . .

RUN gcc src/main.c -o bot -pthread -ldiscord -lcurl

ENTRYPOINT [ "./bot" ]
