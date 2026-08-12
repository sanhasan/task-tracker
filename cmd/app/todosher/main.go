package main

import (
	"fmt"
	"os"
	"task-tracker/internal/task/app"
)

func main() {
	resp, err := app.Run(os.Args[1:])

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("sucsesful: %s\n", resp)
}
