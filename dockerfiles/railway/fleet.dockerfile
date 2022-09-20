FROM botwayorg/fleet-rs:latest

COPY . .

RUN mkdir /root/.botway

COPY botway.json /root/.botway

RUN apk update

# To add more packages
# RUN apk add PACKAGE_NAME

RUN fleet build --release --bin bot

ENTRYPOINT ["./target/release/bot"]
