FROM botwayorg/concord

COPY . .

ENV BOTWAY-DIR /root/.botway

RUN mkdir ${BOTWAY-DIR}

COPY botway.json ${BOTWAY-DIR}

# To add more packages
# RUN apk add PACKAGE_NAME

RUN gcc src/main.c -o bot -pthread -ldiscord -lcurl

ENTRYPOINT [ "./bot" ]
