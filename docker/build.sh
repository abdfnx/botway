#!/bin/bash

tag="$1"
filename=$tag

if [ $tag == "latest" ]; then
    filename="alpine"
fi

docker build -t botwayorg/botway:$tag --file ./docker/$tag.dockerfile .
docker push botwayorg/botway:$tag
