FROM gitpod/workspace-full

# Install task, gomo, deno, and github cli
RUN brew install go-task/tap/go-task abdfnx/tap/gomo deno gh

# Install pnpm
RUN npm i -g npm@latest pnpm@latest

# Install Botway CLI
RUN curl -sL https://dub.sh/botway | bash

# Initialize `~/.botway`
RUN botway init --no-repo
