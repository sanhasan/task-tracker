package main

import (
	"fmt"
	"os"

	"task-tracker/internal/task/app"
	"task-tracker/internal/task/config"
)

func main() {
	cfg := config.New()

	resp, err := app.Run(cfg, os.Args[1:])

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Print(resp)
}
