package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/color"
)

type LogEntry struct {
	DateTime   string `json:"DateTime"`
	LogLevel   string `json:"LogLevel"`
	FileName   string `json:"FileName"`
	MethodName string `json:"MethodName"`
	LineNumber string `json:"LineNumber"`
	LogMessage string `json:"LogMessage"`
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		var entry LogEntry
		err := json.Unmarshal([]byte(scanner.Text()), &entry)
		if err != nil {
			fmt.Println(scanner.Text()) // Print non-JSON lines directly
			continue
		}

		switch entry.LogLevel {
		case "INFO":
			color.New(color.FgGreen).Printf("[%s] %s - %s:%s %s\n", entry.LogLevel, entry.DateTime, entry.FileName, entry.LineNumber, entry.LogMessage)
		case "DEBUG":
			color.New(color.FgYellow).Printf("[%s] %s - %s:%s %s\n", entry.LogLevel, entry.DateTime, entry.FileName, entry.LineNumber, entry.LogMessage)
		case "ERR":
			color.New(color.FgRed).Printf("[%s] %s - %s:%s %s\n", entry.LogLevel, entry.DateTime, entry.FileName, entry.LineNumber, entry.LogMessage)
		default:
			fmt.Printf("[%s] %s - %s:%s %s\n", entry.LogLevel, entry.DateTime, entry.FileName, entry.LineNumber, entry.LogMessage)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
