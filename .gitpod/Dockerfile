FROM gitpod/workspace-full:2023-03-24-22-45-37

# Install task, gomo, deno, and github cli
RUN brew update

RUN brew install go-task/tap/go-task abdfnx/tap/gomo abdfnx/tap/botway deno go gh

# Install pnpm
RUN npm i -g npm@latest pnpm@latest

# Initialize `~/.botway`
RUN botway init --no-repo
