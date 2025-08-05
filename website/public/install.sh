#!/bin/bash

if [[ $EUID -ne 0 ]]; then
    echo "Please run this script as a root user. Use sudo"
    exit 1
else
    
   curl -LO https://github.com/winnerx0/commit-buddy/releases/download/commit-buddy/cb.zip
    
    unzip cb.zip
    
    rm cb.zip
    
    cd cb
    
    current_dir=$(pwd)
    
    go build -o cb .
    
    sudo mv cb /usr/local/bin
    
    cd ..
    
    rm -rf "$current_dir"
    
fi

echo "Installation complete please run cb to get started!"

echo "Please export your Open Router API Key using export OPEN_ROUTER_API_KEY=YOUR_API_KEY"
