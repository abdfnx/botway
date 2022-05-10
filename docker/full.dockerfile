FROM frolvlad/alpine-glibc:alpine-3.14_glibc-2.33

### variables ###
ENV PKGS="zip unzip git curl npm build-base neofetch zsh sudo make lsof wget gcc asciidoctor ca-certificates bash-completion htop jq less llvm nano vim ruby-full ruby-dev libffi-dev"
ENV ZSHRC=".zshrc"

### install packages ###
RUN apk upgrade && \
    apk add --update $PKGS

### setup user ###
USER root
RUN adduser -D bw \
    && echo "bw ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/bw \
    && chmod 0440 /etc/sudoers.d/bw

### nodejs package managers ###
RUN npm i -g npm@latest
RUN npm i -g yarn@latest
RUN npm i -g pnpm@latest

### botway ###
# RUN curl -sL https://botway.web.app/get | bash

ENV HOME="/home/bw"
WORKDIR $HOME
USER bw

### zsh ###
RUN zsh && \
    sh -c "$(curl -fsSL https://raw.github.com/robbyrussell/oh-my-zsh/master/tools/install.sh)" && \
    sudo apk update && \
    git clone https://github.com/zsh-users/zsh-syntax-highlighting ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting && \
    git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions

### rust ###
RUN curl https://sh.rustup.rs -sSf | bash -s -- -y
RUN /bin/bash -c "source ~/.cargo/env"

### go ###
RUN wget "https://dl.google.com/go/$(curl https://go.dev/VERSION?m=text).linux-amd64.tar.gz"
RUN sudo tar -C /usr/local -xzf "$(curl https://go.dev/VERSION?m=text).linux-amd64.tar.gz"
ENV GOROOT /usr/local/go/bin
ENV GOPATH /go
ENV PATH /go/bin:$PATH
RUN rm "$(curl https://go.dev/VERSION?m=text).linux-amd64.tar.gz"

RUN sudo mkdir -p ${GOPATH}/src ${GOPATH}/bin

### deno ###
RUN curl -fsSL https://deno.land/x/install/install.sh | sh

ENV DENO_INSTALL="$HOME/.deno"

ENV PATH="${DENO_INSTALL}/bin:${PATH}"

### rm old ~/.zshrc ###
RUN sudo rm -rf $ZSHRC

COPY ./tools/zshrc $ZSHRC

RUN sudo chown bw $ZSHRC

CMD /bin/bash -c "zsh"
