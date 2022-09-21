FROM python:alpine

COPY . .

RUN mkdir /root/.botway

COPY botway.json /root/.botway

ENV PACKAGES "build-dependencies build-base gcc abuild binutils binutils-doc gcc-doc py-pip jpeg-dev zlib-dev python3 python3-dev libffi-dev git boost boost-dev"

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

# To add more packages
# RUN apk add PACKAGE_NAME

# Install poetry
RUN pip install poetry

RUN poetry config virtualenvs.create false
RUN poetry install

ENTRYPOINT ["python3", "./src/main.py"]
