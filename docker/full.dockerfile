FROM debian:11
FROM botwayorg/botway:debian

### variables ###
ENV UPD="apt-get update"
ENV UPD_s="sudo $UPD"
ENV INS="apt-get install"
ENV INS_s="sudo $INS"
ENV PKGS="zip unzip multitail curl zsh lsof wget ssl-cert asciidoctor apt-transport-https ca-certificates gnupg-agent bash-completion build-essential htop jq software-properties-common less llvm locales man-db nano vim ruby-full"
ENV BUILDS="zlib1g-dev libncurses5-dev libgdbm-dev libnss3-dev libssl-dev libsqlite3-dev libreadline-dev libffi-dev libbz2-dev"
ENV BOTWAY_DOCKER_TOOLS_URL="https://abdfnx.github.io/botway/docker/tools"
ENV LANG=en_US.UTF-8
ENV ZSHRC=".zshrc"

RUN $UPD && $INS -y $PKGS && $UPD && \
    locale-gen en_US.UTF-8 && \
    mkdir /var/lib/apt/abdcodedoc-marks && \
    apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* && \
    $UPD

### git ###
RUN $INS -y git && \
    rm -rf /var/lib/apt/lists/* && \
    $UPD

### sudo ###
RUN $UPD && $INS -y sudo && \
    adduser --disabled-password --gecos '' bw && \
    adduser bw sudo && \
    echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

### nodejs & npm ###
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh | bash
### rm old ~/.zshrc ###
RUN sudo rm -rf $ZSHRC

RUN wget $BOTWAY_DOCKER_TOOLS_URL/zshrc -O $ZSHRC

RUN source $ZSHRC

RUN nvm install 18
RUN nvm alias default 18

ENV HOME="/home/bw"
WORKDIR $HOME
USER bw

### zsh ###
RUN zsh && \
    sh -c "$(curl -fsSL https://raw.github.com/robbyrussell/oh-my-zsh/master/tools/install.sh)" && \
    $UPD_s && \
    git clone https://github.com/zsh-users/zsh-syntax-highlighting ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting && \
    git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions

RUN source $ZSHRC

CMD /bin/bash -c "zsh"
