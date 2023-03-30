FROM botwayorg/gp-image:latest

# Install tools
RUN brew install go-task/tap/go-task abdfnx/tap/gomo abdfnx/tap/botway deno go gh dart-lang/dart/dart

# Install pnpm
RUN npm i -g npm@latest pnpm@latest

# Initialize `~/.botway`
RUN botway init --no-repo
