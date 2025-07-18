package pretty

import (
	"fmt"
	"time"

	"github.com/apex-woot/git-recap/internal/git"
)

type PrettyMsg struct {
	GitCommit *git.GitCommit
	PrettyMsg string
}

type PrettyByDate struct {
	Date       time.Time
	CommitMsgs []*PrettyMsg
}

func PrettifyByDate(commits []*git.GitCommit) map[time.Time][]*PrettyMsg {
	dateMap := map[time.Time][]*PrettyMsg{}

	for _, c := range commits {
		date := time.Date(c.Date.Year(), c.Date.Month(), c.Date.Day(), 0, 0, 0, 0, time.Local)
		messageWithTimestamp := fmt.Sprintf("%s (at %s)", c.Msg, time.Date(0, 0, 0, c.Date.Hour(), c.Date.Minute(), c.Date.Second(), 0, time.Local).Format("15:04:05"))
		pMsg := PrettyMsg{GitCommit: c, PrettyMsg: messageWithTimestamp}

		if _, ok := dateMap[date]; !ok {
			dateMap[date] = []*PrettyMsg{&pMsg}
		} else {
			dateMap[date] = append(dateMap[date], &pMsg)
		}
	}

	return dateMap
}
