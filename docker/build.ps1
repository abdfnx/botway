$tags = "core", "alpine", "alpine-glibc", "centos", "debian", "distroless", "latest", "ubuntu"

foreach ($tag in $tags) {
    $filename = $tag

    if ($tag == "latest") {
        $filename = "alpine"
    }

    docker build -t botwayorg/botway:$tag --file ".\docker\$filename.dockerfile" .
    docker push botwayorg/botway:$tag
}

git clone https://github.com/botwayorg/app-core

cd ./app-core

docker build -t botwayorg/app --build-arg NEXT_PUBLIC_BW_SECRET_KEY=$(echo $Env:BW_SECRET_KEY) .

docker push botwayorg/app

cd ..
