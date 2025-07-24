#!/bin/bash

go build -o cb .

if [[ $EUID -ne 0 ]]; then
    echo "Please run this script as a root user. Use sudo"
    exit 1
else
    mv ./cb /usr/local/bin
fi

echo "Installation complete please run cb to get started!"

echo "Please export your Open Router API Key using export OPEN_ROUTER_API_KEY=YOUR_API_KEY"


