FROM botwayorg/concord

COPY . .

RUN mkdir /root/.botway

COPY botway.json /root/.botway

# To add more packages
# RUN apk add PACKAGE_NAME

RUN gcc src/main.c -o bot -pthread -ldiscord -lcurl

ENTRYPOINT [ "./bot" ]
