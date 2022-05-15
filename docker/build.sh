#!/bin/bash

tags=( alpine alpine-glibc centos core debian distroless latest ubuntu )

for t in "${tags[@]}"
do
    filename=${t}

    if [ "$t" == "latest" ]; then
        filename="alpine"
    fi

	docker build -t botwayorg/botway:$t --file ./docker/$filename.dockerfile .
    docker push botwayorg/botway:$t
done
