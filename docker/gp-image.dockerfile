FROM gitpod/workspace-base:2023-11-04-12-07-48

USER gitpod

RUN sudo apt-get update -yq && sudo apt-get upgrade -y

RUN sudo apt-get install -y git-core curl gnupg build-essential openssl libssl-dev ruby ruby-dev 

### Nodejs ###
RUN curl -fsSL https://deb.nodesource.com/setup_current.x | sudo -E bash - && \
    sudo apt-get install -y nodejs

### Homebrew ###
RUN mkdir ~/.cache && /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
ENV PATH=/home/linuxbrew/.linuxbrew/bin:/home/linuxbrew/.linuxbrew/sbin/:$PATH
ENV MANPATH="$MANPATH:/home/linuxbrew/.linuxbrew/share/man"
ENV INFOPATH="$INFOPATH:/home/linuxbrew/.linuxbrew/share/info"
ENV HOMEBREW_NO_AUTO_UPDATE=1

RUN brew install cmake

### C ###
RUN curl -fsSL https://apt.llvm.org/llvm-snapshot.gpg.key | sudo gpg --dearmor -o /usr/share/keyrings/llvm-archive-keyring.gpg \
    && echo "deb [signed-by=/usr/share/keyrings/llvm-archive-keyring.gpg] http://apt.llvm.org/jammy/ \
    llvm-toolchain-jammy-15 main" | sudo tee /etc/apt/sources.list.d/llvm.list > /dev/null \
    && sudo apt update \
    && sudo install-packages \
    clang \
    clangd \
    clang-format \
    clang-tidy \
    gdb \
    lld

### Rust ###
ENV PATH=$HOME/.cargo/bin:$PATH

RUN curl -fsSL https://sh.rustup.rs | sh -s -- -y --profile minimal --no-modify-path --default-toolchain stable \
    -c rls rust-analysis rust-src rustfmt clippy \
    && for cmp in rustup cargo; do rustup completions bash "$cmp" > "$HOME/.local/share/bash-completion/completions/$cmp"; done \
    && printf '%s\n'    'export CARGO_HOME=/workspace/.cargo' \
    'mkdir -m 0755 -p "$CARGO_HOME/bin" 2>/dev/null' \
    'export PATH=$CARGO_HOME/bin:$PATH' \
    'test ! -e "$CARGO_HOME/bin/rustup" && mv "$(command -v rustup)" "$CARGO_HOME/bin"' > $HOME/.bashrc.d/80-rust \
    && cargo install cargo-watch cargo-edit cargo-workspaces \
    && rm -rf "$HOME/.cargo/registry"

# Install tools using homebrew
RUN brew update
