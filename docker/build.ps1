$tags = "core", "alpine", "alpine-glibc", "centos", "debian", "distroless", "latest", "ubuntu"

foreach ($tag in $tags) {
    $filename = $tag

    if ($tag == "latest") {
        $filename = "alpine"
    }

    docker build -t botwayorg/botway:$tag --file ".\docker\$filename.dockerfile" .
    docker push botwayorg/botway:$tag
}
