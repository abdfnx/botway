#!/bin/bash

tags=( core alpine alpine-glibc centos debian distroless latest ubuntu )

for t in "${tags[@]}"
do
    filename=${t}

    if [ "$t" == "latest" ]; then
        filename="alpine"
    fi

	docker build -t botwayorg/botway-cli:$t --file ./docker/$filename.dockerfile .
    docker push botwayorg/botway-cli:$t
done

docker build -t botwayorg/gp-image --file ./docker/gp-image.dockerfile .
docker push botwayorg/gp-image

cd ./core

docker build -t botwayorg/botway --build-arg NEXT_PUBLIC_BW_SECRET_KEY=$(echo $BW_SECRET_KEY) .

docker push botwayorg/botway

cd ..
