FROM frolvlad/alpine-glibc:alpine-3.14_glibc-2.33

### variables ###
ENV PKGS="zip unzip git curl npm py3-pip openssl openssl-dev libsodium ffmpeg lld clang build-base abuild binutils opus autoconf automake libtool m4 youtube-dl binutils-doc gcc-doc python3-dev neofetch zsh sudo make lsof wget gcc asciidoctor ca-certificates bash-completion htop jq less llvm nano vim ruby-full ruby-dev libffi-dev icu-libs krb5-libs libgcc libintl libssl1.1 libstdc++ zlib"
ENV ZSHRC=".zshrc"

### install packages ###
RUN apk upgrade && \
    apk add --update $PKGS

ENV RUSTUP_HOME=/usr/local/rustup \
    CARGO_HOME=/usr/local/cargo \
    PATH=/usr/local/cargo/bin:$PATH

RUN set -eux; \
    apkArch="$(apk --print-arch)"; \
    case "$apkArch" in \
    x86_64) rustArch='x86_64-unknown-linux-musl' ;; \
    aarch64) rustArch='aarch64-unknown-linux-musl' ;; \
    *) echo >&2 "unsupported architecture: $apkArch"; exit 1 ;; \
    esac; \
    \
    url="https://static.rust-lang.org/rustup/dist/${rustArch}/rustup-init"; \
    wget "$url"; \
    chmod +x rustup-init; \
    ./rustup-init -y --no-modify-path --default-toolchain nightly; \
    rm rustup-init; \
    chmod -R a+w $RUSTUP_HOME $CARGO_HOME; \
    rustup --version; \
    cargo --version; \
    rustc --version;

### setup user ###
USER root
RUN adduser -D bw \
    && echo "bw ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/bw \
    && chmod 0440 /etc/sudoers.d/bw

### nodejs package managers ###
RUN npm i -g npm@latest yarn@latest pnpm@latest

### botway ###
RUN curl -sL dub.sh/botway | bash

ENV HOME="/home/bw"
WORKDIR $HOME
USER bw

### zsh ###
RUN zsh && \
    sh -c "$(curl -fsSL https://raw.github.com/robbyrussell/oh-my-zsh/master/tools/install.sh)" && \
    sudo apk update && \
    git clone https://github.com/zsh-users/zsh-syntax-highlighting ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting && \
    git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions

### update bundler ###
RUN sudo gem update bundler

### go ###
RUN wget "https://dl.google.com/go/$(curl https://go.dev/VERSION?m=text).linux-amd64.tar.gz"
RUN sudo tar -C /usr/local -xzf "$(curl https://go.dev/VERSION?m=text).linux-amd64.tar.gz"
ENV GOROOT /usr/local/go/bin
ENV PATH /go/bin:$PATH
RUN rm "$(curl https://go.dev/VERSION?m=text).linux-amd64.tar.gz"

### deno ###
RUN curl -fsSL https://deno.land/install.sh | sh
ENV DENO_INSTALL="$HOME/.deno"
ENV PATH="${DENO_INSTALL}/bin:${PATH}"

### gh ###
RUN wget \
    https://github.com/cli/cli/releases/download/$(curl https://get-latest.deno.dev/cli/cli)/gh_$(curl https://get-latest.deno.dev/cli/cli?no-v=true)_linux_amd64.tar.gz \
    -O gh.tar.gz
RUN tar -xzf gh.tar.gz
RUN sudo mv "gh_$(curl https://get-latest.deno.dev/cli/cli?no-v=true)_linux_amd64/bin/gh" /usr/bin
RUN rm -rf gh*

### pyenv ###
RUN pip install tld --ignore-installed six distlib --user
RUN curl https://raw.githubusercontent.com/pyenv/pyenv-installer/master/bin/pyenv-installer | bash
# these need to go into your .bashrc
ENV PATH="$HOME/.pyenv/bin:$PATH"
RUN echo 'eval "$(pyenv init -)"' >> ~/.bashrc
RUN echo 'eval "$(pyenv virtualenv-init -)"' >> ~/.bashrc
RUN /bin/bash -c "bash"

### pipenv ###
RUN curl https://raw.githubusercontent.com/pypa/pipenv/master/get-pipenv.py | python3

### poetry ###
RUN curl -sSL https://raw.githubusercontent.com/python-poetry/poetry/master/get-poetry.py | python3

ENV PATH="/root/.poetry/bin:$PATH"
RUN echo 'eval "$(poetry env install -q)"' >> ~/.bashrc
RUN echo 'eval "$(poetry env shell -q)"' >> ~/.bashrc
RUN /bin/bash -c "bash"

### c# ###
RUN curl -sL https://dot.net/v1/dotnet-install.sh | bash

### rm old ~/.zshrc ###
RUN sudo rm -rf $ZSHRC
COPY ./tools/.zshrc .
RUN sudo chown bw $ZSHRC

CMD /bin/bash -c "zsh"
