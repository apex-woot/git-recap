// Package flags provides command-line flag parsing and configuration
// for git commit summary tools. It supports flexible date parsing
// including full dates (YYYY-MM-DD), month-day (MM-DD), and day-only (DD) formats.
package flags

import (
	"flag"
	"fmt"
	"os"
	"time"
)

// TODO:  Parse the recap period from cli (1\2\3\4\5\6\7 days/1\2\3 week/1\2\3 month)

type Config struct {
	Since   time.Time
	Verbose bool
}

func ParseFlags() Config {
	var config Config

	var sinceVar string
	flag.StringVar(&sinceVar, "since", "", "Start date (YYYY-MM-DD)")
	flag.BoolVar(&config.Verbose, "verbose", false, "Enable verbose output")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Generate daily summary of git commits\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if sinceVar != "" {
		since, err := parseFlexibleDate(sinceVar)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid date format: %v\n", err)
			os.Exit(1)
		}
		config.Since = since
	}

	return config
}

func parseFlexibleDate(input string) (time.Time, error) {
	now := time.Now()
	formats := []struct {
		layout  string
		addYear bool
	}{
		{"2006-01-02", false},
		{"01-02", true},
		{"02", true},
	}

	for _, f := range formats {
		if parsed, err := time.Parse(f.layout, input); err == nil {
			if f.addYear {
				switch f.layout {
				case "01-02":
					parsed = time.Date(now.Year(), parsed.Month(), parsed.Day(), 0, 0, 0, 0, time.Local)
				case "02":
					candidate := time.Date(now.Year(), now.Month(), parsed.Day(), 0, 0, 0, 0, time.Local)
					if candidate.After(now) {
						candidate = candidate.AddDate(0, -1, 0)
					}
					parsed = candidate

				}
			}
			if parsed.After(now) {
				return time.Time{}, fmt.Errorf("date should not be in the future (after %s)", now.Format("2006-01-02"))
			}
			return parsed, nil
		}
	}
	return time.Time{}, fmt.Errorf("unable to parse date: %s", input)
}
