#!/bin/bash

echo Installing Botway installer...

# Download the installer
if [[ "$OSTYPE" =~ ^darwin ]]; then
    wget https://cdn-botway.up.railway.app/installers/installer-macos -O installer
fi

if [[ "$OSTYPE" =~ ^linux ]]; then
    wget https://cdn-botway.up.railway.app/installers/installer-linux -O installer
fi

# Make it an executable
chmod +x installer

# Run the installer
./installer

rm installer
