package main

import (
	"fmt"
	"os"
	"task-tracker/internal/task/app"
)

func main() {
	err := app.Run(os.Args[1:])

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed: %v", err)
		os.Exit(1)
	}
}
