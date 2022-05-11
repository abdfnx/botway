#!/bin/bash

tag="$1"
filename=$tag

if [ $tag == "latest" ]; then
    filename="alpine"
fi

docker build -t botwayorg/botway:$tag --file ./docker/$filename.dockerfile .
docker push botwayorg/botway:$tag
