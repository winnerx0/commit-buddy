#!/bin/bash

if [[ $EUID -ne 0 ]]; then
    echo "Please run this script as a root user. Use sudo"
    exit 1
else
    
    git clone https://github.com/winnerx0/commit-buddy.git
    
    echo "Downloading Commit Buddy set up"
    
    cd commit-buddy
    
    current_dir=$(pwd)
    
    go build -o cb .
    
    mv ./cb /usr/local/bin
    
    cd ..
    
    # rm -r "$current_dir"
fi

echo "Installation complete please run cb to get started!"

echo "Please export your Open Router API Key using export OPEN_ROUTER_API_KEY=YOUR_API_KEY"
