ARG VARIANT=1.18
ARG NODE_VERSION="18"

FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

RUN export DEBIAN_FRONTEND=noninteractive && \
    apt-get update -y && \
    apt-get upgrade -y && \
    apt-get install software-properties-common curl -y

# Install Node.js
SHELL ["/bin/bash", "--login", "-i", "-c"]

RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.2/install.sh | bash

RUN source /root/.bashrc && nvm install 18

SHELL ["/bin/bash", "--login", "-c"]

# Install pnpm
RUN npm i -g npm@latest pnpm@latest

# Install Botway CLI
RUN curl -sL https://dub.sh/botway | bash

# Initialize `~/.botway`
RUN botway init --no-repo

# Install Homebrew
RUN git clone https://github.com/Homebrew/brew ~/.linuxbrew/Homebrew \
    && mkdir ~/.linuxbrew/bin \
    && ln -s ../Homebrew/bin/brew ~/.linuxbrew/bin \
    && echo 'eval $(~/.linuxbrew/bin/brew shellenv)' >> /root/.bashrc && \
    source /root/.bashrc

# Install deno and go-task
RUN brew install deno go-task/tap/go-task
