FROM botwayorg/botway:latest AS bw

COPY . .

RUN botway init --docker

### Add here the linux distribution you want to use ###
### Example: FROM alpine:latest ###
### Or with with the bot language with the tag: ###
### Example: FROM ruby:alpine ###

### Add here the packages you want to use ###
### Example: ENV PACKAGES "Packages to add" ###
### RUN apk update && apk add --no-cache --virtual ${PACKAGES} ###

### Copy botway config from bw to /root/.botway ###
### Example: COPY --from=bw /root/.botway /root/.botway ###

### Copy here all the files you want to use ###
### Example: COPY . . ### 

### Add here the build command to build your bot ###
### Example: RUN cargo build --release ###

### Last step: Add here the entrypoint command to run your bot ###
### Example: ENTRYPOINT ["python3", "./src/main.py"] ###
