docker build -t botwayorg/botway:$1 --file ./docker/$1.dockerfile .
docker push botwayorg/botway:$1
