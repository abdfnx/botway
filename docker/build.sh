#!/bin/bash

tag="$1"

if [ $tag == "latest" ]; then
    tag="alpine"
fi

docker build -t botwayorg/botway:$tag --file ./docker/$tag.dockerfile .
docker push botwayorg/botway:$tag
