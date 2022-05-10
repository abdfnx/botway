export ZSH="/home/bw/.oh-my-zsh"

ZSH_THEME="af-magic"

plugins=( git zsh-syntax-highlighting zsh-autosuggestions )

source $ZSH/oh-my-zsh.sh

alias s="source ~/.zshrc"
alias n="nano ~/.zshrc"
alias update="sudo apt-get update"
alias upgrade="sudo apt-get upgrade"
alias py="python3"
alias go="/usr/local/go/bin/go"
