package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/apex-woot/git-recap/internal/flags"
)

func main() {
	// Parse the recap period from cli (since date)
	cfg := flags.ParseFlags()
	fmt.Printf("%+v\n\n", cfg)

	cmd := exec.Command("git", "log", "--oneline")
	output, err := cmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute git log command: %v", err)
	}

	fmt.Println(output)
	// Check if the folder is a root repo folder (has .git) folder
	// Get commit history using git api
	// Parse the history and map days-to-group of commits done in that day
	// Output the history to the io.Writer
}
