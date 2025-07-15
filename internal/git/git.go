package git

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type GitCommit struct {
	ShortHash string
	Msg       string
	Author    string
	Date      time.Time
}

func GetCommits(since time.Time) ([]*GitCommit, error) {
	cmd := exec.Command("git", "log", "--pretty=format:%h|%s|%an|%ad", "--date=iso", fmt.Sprintf("--since=%s", since.Format("2006-01-02")))
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to execute git log command: %v", err)
	}
	commits := strings.FieldsFunc(string(output), func(r rune) bool { return r == '\n' })
	gitCommits := make([]*GitCommit, len(commits))
	for i, c := range commits {
		gitCommitParts := strings.FieldsFunc(strings.TrimSpace(c), func(r rune) bool { return r == '|' })
		timestamp, err := time.Parse("2006-01-02 15:04:05 -0700", gitCommitParts[3])
		if err != nil {
			return nil, fmt.Errorf("can not parse commit date: %w", err)
		}
		gitCommit := GitCommit{ShortHash: gitCommitParts[0], Msg: gitCommitParts[1], Author: gitCommitParts[2], Date: timestamp}
		gitCommits[i] = &gitCommit
	}

	for _, c := range gitCommits {
		fmt.Fprintf(os.Stderr, "%+v\n", *c)
	}
	return gitCommits, nil
}
