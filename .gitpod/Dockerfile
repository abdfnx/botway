FROM denoland/deno:latest AS deno
FROM gitpod/workspace-full

COPY --from=deno /deno /usr/local/bin/deno

# Install task, gomo, deno, and github cli
RUN brew install go-task/tap/go-task abdfnx/tap/gomo abdfnx/tap/botway gh

# Install pnpm
RUN npm i -g npm@latest pnpm@latest

# Initialize `~/.botway`
RUN botway init --no-repo
