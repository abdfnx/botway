FROM gitpod/workspace-full

# Install task, gomo, deno, and github cli
RUN brew update && brew install go-task/tap/go-task abdfnx/tap/gomo abdfnx/tap/botway deno go gh

# Install pnpm
RUN npm i -g npm@latest pnpm@latest

# Initialize `~/.botway`
RUN botway init --no-repo
