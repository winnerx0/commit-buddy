package main

import (
	"fmt"
	"os/exec"
)

func main(){



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

	fmt.Println(string(diffOutput))
}