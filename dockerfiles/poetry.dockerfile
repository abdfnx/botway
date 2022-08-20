FROM python:alpine
FROM botwayorg/botway:latest

ENV PACKAGES "build-dependencies build-base gcc abuild binutils binutils-doc gcc-doc py-pip jpeg-dev zlib-dev python3 python3-dev libffi-dev git boost boost-dev"

COPY . .

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

# Add packages you want
# RUN apk add PACKAGE_NAME

# Install poetry
RUN curl -sSL https://raw.githubusercontent.com/python-poetry/poetry/master/get-poetry.py | python3

ENV PATH="/root/.poetry/bin:$PATH"

RUN echo 'eval "$(poetry env install -q)"' >> ~/.bashrc
RUN echo 'eval "$(poetry env shell -q)"' >> ~/.bashrc

RUN /bin/bash -c "bash"

RUN poetry config virtualenvs.create false
RUN poetry install --no-dev
