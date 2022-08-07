#!/bin/bash

echo Installing Botway installer...

# Download the installer
if [[ "$OSTYPE" =~ ^darwin ]]; then
    wget https://cdn-botway.waypoint.run/installers/installer-macos -O bw-installer
fi

if [[ "$OSTYPE" =~ ^linux ]]; then
    wget https://cdn-botway.waypoint.run/installers/installer-linux -O bw-installer
fi

# Make it an executable
chmod +x bw-installer

# Run the installer
./bw-installer

rm bw-installer
