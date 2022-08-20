FROM botwayorg/botway:latest AS bw

COPY . .

RUN botway init --docker

FROM botwayorg/fleet-rs:latest

RUN apk update

# Add packages you want
# RUN apk add PACKAGE_NAME

COPY --from=bw /root/.botway /root/.botway

COPY . .

RUN fleet build --release --bin bot

ENTRYPOINT ["./target/release/bot"]
