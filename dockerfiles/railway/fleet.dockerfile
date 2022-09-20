FROM botwayorg/fleet-rs:latest

COPY . .

ENV BOTWAY-DIR /root/.botway

RUN mkdir ${BOTWAY-DIR}

COPY botway.json ${BOTWAY-DIR}

RUN apk update

# To add more packages
# RUN apk add PACKAGE_NAME

RUN fleet build --release --bin bot

ENTRYPOINT ["./target/release/bot"]
