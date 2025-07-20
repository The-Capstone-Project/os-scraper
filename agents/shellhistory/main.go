package shellhistory

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func Run() {
	// Set default limit to 10
	n := flag.Int("n", 10, "Number of history items to fetch (default: 10)")
	flag.Parse()
	shellpath := os.Getenv("SHELL")
	shellarr := strings.Split(string(shellpath), "/")
	shell := shellarr[len(shellarr)-1]
	var HistoryFile string

	if shell == "zsh" {
		HistoryFile = os.Getenv("HOME") + "/.zsh_history"
	} else if shell == "bash" {
		HistoryFile = os.Getenv("HOME") + "/.bash_history"
	} else if shell == "fish" {
		HistoryFile = os.Getenv("HOME") + "/.local/share/fish/fish_history"
	}

	// If n is <= 0, default to 10
	if *n <= 0 {
		*n = 10
	}

	printLimitedHistory(HistoryFile, *n, shell)
}

func printLimitedHistory(filename string, n int, shell string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading %s shell history file:\n%s\n", shell, err)
		return
	}
	lines := strings.Split(string(content), "\n")
	start := len(lines) - n
	if start < 0 {
		start = 0
	}
	fmt.Printf("Last %d lines of %s history file:\n", n, shell)
	for _, line := range lines[start:] {
		fmt.Println(line)
	}
}
