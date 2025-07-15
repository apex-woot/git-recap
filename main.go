package main

import (
	"fmt"
	"log"

	"github.com/apex-woot/git-recap/internal/flags"
	"github.com/apex-woot/git-recap/internal/git"
)

func main() {
	// Parse the recap period from cli (since date)
	cfg := flags.ParseFlags()
	fmt.Printf("%+v\n\n", cfg)

	// Check if the folder is a root repo folder (has .git) folder

	// Get commit history using git api
	_, err := git.GetCommits(cfg.Since)
	if err != nil {
		log.Fatal(err)
	}
	// Parse the history and map days-to-group of commits done in that day
	// Output the history to the io.Writer
}
