FROM alpine:latest
FROM botwayorg/botway:alpine

### variables ###
ENV PKGS="zip unzip git curl npm build-base zsh sudo make go lsof wget gcc asciidoctor ca-certificates bash-completion htop jq less llvm man-db nano vim ruby-full ruby-dev libffi-dev"
ENV ZSHRC=".zshrc"

### install packages ###
RUN apk update && \
    apk add $PKGS

### setup user ###
USER root
RUN adduser -D bw \
    && echo "bw ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/bw \
    && chmod 0440 /etc/sudoers.d/bw

### nodejs package managers ###
RUN npm i -g npm@latest
RUN npm i -g yarn@latest
RUN npm i -g pnpm@latest

ENV HOME="/home/bw"
WORKDIR $HOME
USER bw

### zsh ###
RUN zsh && \
    sh -c "$(curl -fsSL https://raw.github.com/robbyrussell/oh-my-zsh/master/tools/install.sh)" && \
    sudo apk update && \
    git clone https://github.com/zsh-users/zsh-syntax-highlighting ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting && \
    git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions

### install rust ###
RUN curl https://sh.rustup.rs -sSf | bash -s -- -y
RUN /bin/bash -c "source ~/.cargo/env"

### go ###
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN sudo mkdir -p ${GOPATH}/src ${GOPATH}/bin

### deno ###
RUN curl -s https://get-latest.herokuapp.com/denoland/deno >> tag.txt
RUN curl -fsSL "https://github.com/denoland/deno/releases/download/$(cat tag.txt)/deno-x86_64-unknown-linux-gnu.zip" \
    --output deno.zip \
  && unzip deno.zip \
  && rm deno.zip \
  && chmod 755 deno \
  && sudo mv deno /usr/bin

### rm old ~/.zshrc ###
RUN sudo rm -rf $ZSHRC

COPY ./tools/zshrc $ZSHRC

# RUN sudo chown bw $ZSHRC

CMD /bin/bash -c "zsh"
