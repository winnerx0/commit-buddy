#!/bin/bash
set -e

if [[ $EUID -ne 0 ]]; then
    echo "❌ Please run this script as root. Use sudo."
    exit 1
fi

# Download and unzip
curl -LO https://github.com/winnerx0/commit-buddy/releases/download/1.1/cb.zip

# Get the folder name before unzipping
dir_name=$(unzip -Z1 cb.zip | head -n1 | cut -d/ -f1)
unzip cb.zip
rm cb.zip

# Change to extracted directory
cd "$dir_name"

# Build and move binary
go build -o cb .
mv cb /usr/local/bin

# Cleanup
cd ..
rm -rf "$dir_name"
rm -- "$0" 2>/dev/null || true

# Success message
clear
cat << "EOF"

  ____                          _ _     ____            _     _       
 / ___|___  _ __ ___  _ __ ___ (_) |_  | __ ) _   _  __| | __| |_   _ 
| |   / _ \| '_ ` _ \| '_ ` _ \| | __| |  _ \| | | |/ _` |/ _` | | | |
| |__| (_) | | | | | | | | | | | | |_  | |_) | |_| | (_| | (_| | |_| |
 \____\___/|_| |_| |_|_| |_| |_|_|\__| |____/ \__,_|\__,_|\__,_|\__, |
                                                                |___/ 

             🚀 Commit Buddy Installed 🚀

   📟 Run it now using:
       cb

   🔑 Set your Gemini API key like this:
       export GEMINI_API_KEY=your_key_here

   ✅ Happy committing!

EOF
