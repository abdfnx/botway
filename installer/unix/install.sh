#!/bin/bash

echo Installing Botway installer...

# Download the installer
if [[ "$OSTYPE" =~ ^darwin ]]; then
    wget https://botway.web.app/installer-macos -O installer
fi

if [[ "$OSTYPE" =~ ^linux ]]; then
    wget https://botway.web.app/installer-linux -O installer
fi

# Make it an executable
chmod +x installer

# Run the installer
./installer

rm installer
