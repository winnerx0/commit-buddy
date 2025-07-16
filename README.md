# Commit Buddy

Commit Buddy is an open-source Command Line Interface (CLI) tool written in Go that leverages the Gemini API to automatically generate intelligent Git commit messages based on your staged changes. Say goodbye to generic commit messages and hello to a more descriptive and meaningful commit history!

## ✨ Features

* **AI-Powered Commit Messages**: Generates concise and relevant commit messages by analyzing your `git diff --staged` output using the Gemini API.

* **Seamless CLI Integration**: Integrates directly into your Git workflow, allowing you to generate messages with simple commands.

* **Interactive Commit Prompt**: Displays the generated `git commit -m "message"` command directly in your terminal's input line, allowing you to edit it, press Enter to commit, or Ctrl+C to abort, just like a normal terminal command.

* **Open Source**: Built with transparency and community contributions in mind.

## 🚀 Getting Started

### Prerequisites

Before you begin, ensure you have the following installed:

* **Go (1.16 or higher)**: [Download and Install Go](https://golang.org/dl/)

* **Git**: [Download and Install Git](https://git-scm.com/)

* **Gemini API Key**: You'll need an API key from Google's Gemini API. You can obtain one by following the instructions on the [Google AI Studio](https://aistudio.google.com/) website.

### Installation

```bash

git clone https://github.com/winnerx0/commit-buddy.git

cd commit-buddy

go build -o cb .

```

This will create an executable named `cb` in your current directory.

(Optional) Add to your `PATH`:

```bash

# On Linux/macOS

sudo mv cb /usr/local/bin/

```

### API Key Setup

Commit Buddy requires your Gemini API key to function. It reads this key from an environment variable named `GEMINI_API_KEY`.

**Recommended for current session:**

```bash

export GEMINI_API_KEY="YOUR_GEMINI_API_KEY_HERE"

```

**For persistent setup:**

* **Bash/Zsh**: Add the export line to your `~/.bashrc`, `~/.zshrc`, or `~/.profile` file.

* **Windows (Command Prompt)**:

  ```cmd

  setx GEMINI_API_KEY "YOUR_GEMINI_API_KEY_HERE"

  ```

* **Windows (PowerShell)**:

  ```powershell

  $env:GEMINI_API_KEY="YOUR_GEMINI_API_KEY_HERE"

  # To make it persistent, add this to your PowerShell profile script

  ```

## 💡 Usage

Navigate to your Git repository, stage your changes, and then run Commit Buddy:

```bash

git add .

cb

```

Commit Buddy will then:

1. Fetch the staged changes (`git diff --staged`).

2. Send the diff to the Gemini API.

3. Display the AI-generated commit message.

**Generated Commit Message Example:**

```

---

feat(user): Add new user authentication module

- Introduce a new authentication modul to handle user login and registration.

---

```

The `git commit` command will then appear on your terminal's input line, pre-filled with the generated message. You can now:

* **Press Enter**: To execute the command as is.

* **Press Ctrl+C**: To abort the commit and exit Commit Buddy.

## 🤝 Contributing

We welcome contributions to Commit Buddy! If you have ideas for new features, bug fixes, or improvements, please feel free to:

1. Fork the repository.

2. Create a new branch:

   ```bash

   git checkout -b feature/your-feature-name

   ```

3. Make your changes.

4. Commit your changes:

   ```bash

   git commit -m 'feat: Add new feature'

   ```

5. Push to the branch:

   ```bash

   git push origin feature/your-feature-name

   ```

6. Open a Pull Request.

Please ensure your code adheres to Go best practices and includes appropriate tests.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.Commit Buddy
