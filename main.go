package main

import (
	"fmt"

	"github.com/apex-woot/git-recap/internal/flags"
)

func main() {
	// Parse the recap period from cli (since date)
	cfg := flags.ParseFlags()
	fmt.Printf("%+v\n\n", cfg)

	// Check if the folder is a root repo folder (has .git) folder
	//
	// Get commit history using git api
	// Parse the history and map days-to-group of commits done in that day
	// Output the history to the io.Writer
}
