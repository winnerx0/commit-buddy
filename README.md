# Commit Buddy
Commit Buddy is an open-source Command Line Interface (CLI) tool written in Go that leverages the Gemini API (via OpenRouter) to automatically generate intelligent Git commit messages based on your staged changes. Say goodbye to generic commit messages and hello to a more descriptive and meaningful commit history!
✨ Features
AI-Powered Commit Messages: Generates concise and relevant commit messages by analyzing your git diff --staged output using the Gemini API (via OpenRouter).
Seamless CLI Integration: Integrates directly into your Git workflow, allowing you to generate messages with simple commands.
Interactive Commit Prompt: Displays the generated git commit -m "message" command directly in your terminal's input line, allowing you to edit it, press Ctrl+S to commit, or Ctrl+C to abort.
Open Source: Built with transparency and community contributions in mind.
🚀 Getting Started
Prerequisites
Before you begin, ensure you have the following installed:
Go (1.16 or higher): Download and Install Go
Git: Download and Install Git
OpenRouter API Key: You'll need an API key from OpenRouter. You can obtain one by following the instructions on the OpenRouter website.
Installation
git clone https://github.com/winnerx0/commit-buddy.git
cd commit-buddy
go build -o cb .


This will create an executable named cb in your current directory.
(Optional) Add to your PATH:
# On Linux/macOS
sudo mv cb /usr/local/bin/


API Key Setup
Commit Buddy requires your OpenRouter API key to function. It reads this key from an environment variable named OPEN_ROUTER_API_KEY.
Recommended for current session:
export OPEN_ROUTER_API_KEY="YOUR_OPEN_ROUTER_API_KEY_HERE"


For persistent setup:
Bash/Zsh: Add the export line to your ~/.bashrc, ~/.zshrc, or ~/.profile file.
💡 Usage
Navigate to your Git repository, stage your changes, and then run Commit Buddy:
git add .
cb


Commit Buddy will then:
Fetch the staged changes (git diff --staged).
Send the diff to the OpenRouter API.
Display the AI-generated commit message.
Generated Commit Message Example:
---
feat(user): Add new user authentication module
- Introduce a new authentication module to handle user login and registration.
---


The generated commit message will appear in an interactive text area. You can now:
Press Ctrl+S: To execute the git commit -m "message" command with the displayed message.
Press Ctrl+C: To abort the commit and exit Commit Buddy.
🤝 Contributing
We welcome contributions to Commit Buddy! If you have ideas for new features, bug fixes, or improvements, please feel free to:
Fork the repository.
Create a new branch:
git checkout -b feature/your-feature-name


Make your changes.
Commit your changes:
git commit -m 'feat: Add new feature'


Push to the branch:
git push origin feature/your-feature-name


Open a Pull Request.
Please ensure your code adheres to Go best practices and includes appropriate tests.
📄 License
This project is licensed under the MIT License - see the LICENSE file for details.

