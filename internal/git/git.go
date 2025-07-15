package git

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GetCommits() []string {
	cmd := exec.Command("git", "log", "--oneline")
	output, err := cmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute git log command: %v", err)
	}

	commits := strings.Split(string(output), "\n")
	fmt.Println(commits)
	return commits
}
