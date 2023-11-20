FROM gitpod/workspace-full:latest

USER gitpod

RUN sudo apt-get update -y && sudo apt-get upgrade -y

# Install tools
RUN brew update && brew install cmake go-task/tap/go-task abdfnx/tap/gomo abdfnx/tap/botway deno go gh dart-lang/dart/dart poetry

# Install pnpm
RUN npm i -g npm@latest pnpm@latest

# Initialize `~/.botway`
RUN botway init --no-repo
