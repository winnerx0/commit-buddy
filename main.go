package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
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

type ContentPart struct {
	Text string `json:"text"`
}

type ContentItem struct {
	Parts []ContentPart `json:"parts"`
}

type GenerateReq struct {
	Contents []ContentItem `json:"contents"`
}

type Error struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"error"`
}

func init() {
	godotenv.Load(".env")
}

func main() {

	client := &http.Client{}

	// _ := flag.String("g", "", "Generate a commit")

	// flag.Parse()

	command := exec.Command("git", "diff", "--staged")

	diffOutput, err := command.Output()

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	if len(diffOutput) == 0 {
		fmt.Println("Please run git add before commit buddy")
		return
	}

	var reqBody = GenerateReq{
		Contents: []ContentItem{
			{
				Parts: []ContentPart{
					{
						Text: fmt.Sprintf(
							"Please use '%s' and create professional Git commit messages based on this diff.",
							string(diffOutput),
						),
					},
				},
			},
		},
	}

	bodyBytes, err := json.MarshalIndent(reqBody, "", "  ")
	if err != nil {
		fmt.Printf("failed to encode request body: %v", err)
	}

	req, err := http.NewRequest("POST", "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent", bytes.NewReader([]byte(bodyBytes)))

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	defer req.Body.Close()

	req.Header.Add("X-goog-api-key", os.Getenv("GEMINI_API_KEY"))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error ", err)
		return
	}

	if res.StatusCode != 200 {
		var error Error

		err = json.Unmarshal(body, &error)

		if err != nil {
			fmt.Println("Error parsing", err)
			return
		}
		fmt.Println(error.Error.Message)
		return
	}

	var aiResponse AIResponse

	err = json.Unmarshal(body, &aiResponse)

	if err != nil {
		fmt.Println("Error parsing", err)
		return
	}

	commit := strings.ReplaceAll(aiResponse.Candidates[0].Content.Parts[0].Text, "```", "")
	
	fmt.Println(commit)
	fmt.Print("Press Enter to run this command, or Ctrl+C to cancel...")

	fmt.Scanln()

	cmd := exec.Command("git", "commit", "-m", commit)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

}
