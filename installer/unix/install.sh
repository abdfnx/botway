#!/bin/bash

echo Installing Botway installer...

# Download the installer
if [[ "$OSTYPE" =~ ^darwin ]]; then
    wget https://cli-botway.deno.dev/installers/installer-macos -O bw-installer
fi

if [[ "$OSTYPE" =~ ^linux ]]; then
    wget https://cli-botway.deno.dev/installers/installer-linux -O bw-installer
fi

# Make it an executable
chmod +x bw-installer

# Run the installer
./bw-installer

rm bw-installer
