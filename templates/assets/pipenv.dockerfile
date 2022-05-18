FROM python:alpine
FROM botwayorg/botway:latest

ENV PACKAGES "build-dependencies build-base gcc abuild binutils binutils-doc gcc-doc python2 py-pip jpeg-dev zlib-dev python3 python3-dev libffi-dev git"

COPY . .

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

# Add packages you want
# RUN apk add PACKAGE_NAME

# Install pyenv
RUN pip install tld --ignore-installed six distlib --user
RUN curl https://raw.githubusercontent.com/pyenv/pyenv-installer/master/bin/pyenv-installer | bash
# these need to go into your .bashrc
ENV PATH="/root/.pyenv/bin:$PATH"
RUN echo 'eval "$(pyenv init -)"' >> ~/.bashrc
RUN echo 'eval "$(pyenv virtualenv-init -)"' >> ~/.bashrc
RUN /bin/bash -c "bash"

# Install pipenv and deps
RUN botway init --docker
RUN curl https://raw.githubusercontent.com/pypa/pipenv/master/get-pipenv.py | python3
RUN pipenv lock
RUN pipenv sync --system
RUN pipenv install

EXPOSE 8000

ENTRYPOINT ["pipenv", "run", "python3", "./src/main.py"]
