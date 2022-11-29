ARG VARIANT=1.18
ARG NODE_VERSION="18"

FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

RUN export DEBIAN_FRONTEND=noninteractive && \
    apt-get update -y && \
    apt-get upgrade -y && \
    apt-get install software-properties-common curl -y

# Install Node.js
RUN if [ "${NODE_VERSION}" != "none" ]; then su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

# Install pnpm
RUN su userx -c "npm i -g pnpm"

# Install Botway CLI
RUN su userx -c "curl -sL https://dub.sh/botway | bash"

# Initialize `~/.botway`
RUN botway init --no-repo

# Install Deps
RUN pnpm i

# Install go-task
RUN go install github.com/go-task/task/v3/cmd/task@latest
