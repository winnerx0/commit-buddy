package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type AIResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GenerateReq struct {
	Contents []Content `json:"contents"`
}

type Content struct {
	Parts []Part `json:"parts"`
}

type Part struct {
	Text string `json:"text"`
}

type Text struct {
	Text string
}
type Error struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"error"`
}

type model struct {
	textarea    textarea.Model
	spinner     spinner.Model
	quitting    bool
	err         error
	commit      string
	spinning    bool
	waiting     bool
	initialized bool
	editing     bool
}

type errMsg struct {
	err error
}

type commitMsg struct {
	commit string
}

type commitDoneMsg struct {
	err error
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick, generateCommit(), textarea.Blink,
	)
}

func initialModel() model {
	ta := textarea.New()
	ta.Placeholder = "Write your commit message..."
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("260"))
	return model{
		spinner:  s,
		spinning: true,
		textarea: ta,
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.textarea, cmd = m.textarea.Update(msg)
	m.spinner, cmd = m.spinner.Update(msg)
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.textarea.SetWidth(msg.Width)
		m.textarea.SetHeight(strings.Count(m.commit, "\n") + 3)
		m.initialized = true

		return m, nil
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlS && m.commit != "" {
			m.editing = false
			m.commit = m.textarea.Value()
			return m, commitCode(m.commit)
		}

		switch msg.String() {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		case "e":
			if m.waiting {

				m.textarea.SetHeight(strings.Count(m.commit, "\n") + 3)
				m.editing = true
				m.waiting = false
				m.textarea.Focus()
				m.textarea.SetValue(m.commit)
				return m, cmd
			}
			return m, nil
		case "enter":
			if m.commit != "" && m.waiting {
				m.editing = false
				return m, commitCode(m.commit)
			}
		default:
			return m, nil
		}
	case errMsg:
		m.err = msg.err
		m.spinning = false
		m.quitting = true
		return m, nil
	case commitMsg:
		var cmd tea.Cmd
		m.spinning = false
		m.commit = msg.commit
		m.waiting = true
		return m, cmd
	case commitDoneMsg:
		return m, tea.Quit

	}

	return m, cmd
}

func (m model) View() string {
	if os.Getenv("GEMINI_API_KEY") == "" {
		return fmt.Sprintf("\n %s", "Please export your Gemini Api Key")
	}

	if m.err != nil {
		return fmt.Sprintf("\n %s", m.err.Error())
	}

	if m.spinning {
		return fmt.Sprintf("\n %s Generating Commit...\n", m.spinner.View())
	}

	if m.commit != "" && m.waiting {
		return fmt.Sprintf("%s\n\n Press Enter to Commit or e to Edit or Ctrl+C to Cancel", m.commit)
	}

	if m.commit != "" && !m.waiting && m.editing {
		return fmt.Sprintf("%s\n Press Ctrl+S to Commit or Ctrl+C to Cancel", m.textarea.View())
	}

	return ""
}

func main() {
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}
}

func isGitRepo() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	err := cmd.Run()
	return err == nil
}

func generateCommit() tea.Cmd {
	return func() tea.Msg {

		if !isGitRepo() {
			return errMsg{err: errors.New("Not inside a Git repository. Run `git init` first.")}
		}

		client := &http.Client{}

		c := exec.Command("git", "diff", "--staged")

		diffOutput, err := c.Output()
		if err != nil {
			return errMsg{err: err}
		}

		if len(diffOutput) <= 0 {
			return errMsg{err: errors.New("No staged changes found. Please stage changes using `git add .` or `git add <file>`")}
		}

		reqBody := GenerateReq{
			Contents: []Content{
				{
					Parts: []Part{
						{
							Text: fmt.Sprintf(
								"You are an AI commit assistant. Based on the following Git diff, generate a high-quality, conventional commit message with the following structure:\n\n1. A single-line header:\n   <type>(<scope>): <short summary>\n   - Use a valid conventional commit type (e.g., feat, fix, refactor, docs, test, chore, style, ci)\n   - Write the summary in the imperative mood (e.g., 'add support for X')\n\n2. A bullet point list describing the main technical changes:\n   - Mention key files, components, classes, or functions changed or added\n   - Use inline code formatting for file names and class/function names (e.g., `someFile.js`, `SomeClass`)\n   - Explain each item concisely and clearly\n\nExample output:\n\n<type>: <short, clear summary of the change>\n- Added SomeUtility to handle core logic for X\n- Updated SomeComponent to support new behavior Y\n- Refactored someFile.js for improved performance\n\nOnly return the non formatted message — no extra explanation or commentary. If you are not confident about a message or what something does **strictly** do not add it to the commit message\n\nGit diff:\n\n%s ",
								string(diffOutput)),
						},
					},
				},
			},
		}

		bodyBytes, err := json.MarshalIndent(reqBody, "", "  ")
		if err != nil {
			return errMsg{err: errors.New("Failed to encode request body")}
		}

		req, err := http.NewRequest("POST", "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent", bytes.NewReader([]byte(bodyBytes)))
		if err != nil {
			return errMsg{err: err}
		}

		req.Header.Add("X-goog-api-key", os.Getenv("GEMINI_API_KEY"))
		req.Header.Add("Content-Type", "application/json")

		res, err := client.Do(req)
		if err != nil {
			return errMsg{err: err}
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return errMsg{err: err}
		}

		if res.StatusCode != 200 {
			var error Error

			err = json.Unmarshal(body, &error)
			if err != nil {
				return errMsg{err: err}
			}
		}

		defer res.Body.Close()

		var aiResponse AIResponse

		err = json.Unmarshal(body, &aiResponse)

		if err != nil {
			return errMsg{err: err}
		}

		if len(aiResponse.Candidates) == 0 {
			return errMsg{err: errors.New("No commit generated")}
		}

		commit := strings.ReplaceAll(aiResponse.Candidates[0].Content.Parts[0].Text, "```", "")

		return commitMsg{
			commit: commit,
		}
	}
}

func commitCode(commit string) tea.Cmd {
	return func() tea.Msg {
		tmpFile, err := os.CreateTemp("", "commit-msg-*.txt")
		if err != nil {
			return commitDoneMsg{err}
		}
		defer os.Remove(tmpFile.Name())

		_, err = tmpFile.WriteString(commit)
		if err != nil {
			return commitDoneMsg{err}
		}
		tmpFile.Close()

		c := exec.Command("git", "commit", "-F", tmpFile.Name())
		err = c.Run()
		return commitDoneMsg{err}
	}
}
