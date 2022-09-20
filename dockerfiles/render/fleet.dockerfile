FROM botwayorg/botway:latest AS bw

COPY . .

RUN botway docker-init

FROM botwayorg/fleet-rs:latest

RUN apk update

# To add more packages
# RUN apk add PACKAGE_NAME

COPY --from=bw /root/.botway /root/.botway

COPY . .

RUN fleet build --release --bin bot

ENTRYPOINT ["./target/release/bot"]
