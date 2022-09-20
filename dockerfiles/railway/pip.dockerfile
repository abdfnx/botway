FROM python:alpine

COPY . .

RUN mkdir /root/.botway

COPY botway.json /root/.botway

ENV PACKAGES "build-dependencies build-base gcc abuild binutils py-pip binutils-doc gcc-doc python3-dev libffi-dev git binutils openssl-dev zlib-dev boost boost-dev"

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

# To add more packages
# RUN apk add PACKAGE_NAME

RUN pip3 install -r requirements.txt

ENTRYPOINT ["python3", "./src/main.py"]
