FROM botwayorg/gp-image:latest

USER gitpod

# Install tools
RUN brew install go-task/tap/go-task abdfnx/tap/gomo abdfnx/tap/botway deno go gh dart-lang/dart/dart poetry

# Install pnpm
RUN sudo npm i -g npm@latest pnpm@latest

# Initialize `~/.botway`
RUN botway init --no-repo
