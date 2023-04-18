FROM golang:alpine

### install packages ###
ENV PKGS "zip unzip git curl npm py3-pip openssl openssl-dev build-base autoconf automake libtool gcc-doc python3-dev neofetch make wget gcc ca-certificates llvm nano vim ruby-full ruby-dev libffi-dev libgcc libssl1.1 zlib"

RUN apk upgrade && \
    apk add --update $PKGS

### github cli ###
RUN wget \
    https://github.com/cli/cli/releases/download/$(curl https://get-latest.deno.dev/cli/cli)/gh_$(curl https://get-latest.deno.dev/cli/cli?no-v=true)_linux_amd64.tar.gz \
    -O gh.tar.gz
RUN tar -xzf gh.tar.gz
RUN mv "gh_$(curl https://get-latest.deno.dev/cli/cli?no-v=true)_linux_amd64/bin/gh" /usr/bin
RUN rm -rf gh*

### install nodejs package managers and create-botway-bot ###
RUN npm i -g npm@latest yarn@latest pnpm@latest create-botway-bot@latest

### update bundler using gem ###
RUN gem update bundler

### pyenv ###
RUN pip install tld --ignore-installed six distlib --user
RUN curl https://raw.githubusercontent.com/pyenv/pyenv-installer/master/bin/pyenv-installer | bash

ENV PATH "$HOME/.pyenv/bin:$PATH"

RUN echo 'eval "$(pyenv init -)"' >> ~/.bashrc
RUN echo 'eval "$(pyenv virtualenv-init -)"' >> ~/.bashrc

RUN /bin/bash -c "bash"

### pipenv ###
RUN curl https://raw.githubusercontent.com/pypa/pipenv/master/get-pipenv.py | python3

### poetry ###
RUN curl -sSL https://install.python-poetry.org | python3 -

ENV PATH "/root/.poetry/bin:$PATH"

RUN echo 'eval "$(poetry env install -q)"' >> ~/.bashrc
RUN echo 'eval "$(poetry env shell -q)"' >> ~/.bashrc

RUN /bin/bash -c "bash"

ARG NEXT_PUBLIC_BW_SECRET_KEY NEXT_PUBLIC_SUPABASE_URL NEXT_PUBLIC_SUPABASE_ANON_KEY

### build ###
RUN git clone https://github.com/abdfnx/botway && \
    cp botway/core .

RUN pnpm i

ENV PORT 3000

EXPOSE 3000

ENTRYPOINT [ "pnpm", "dev" ]
