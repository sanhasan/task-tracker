package main

import (
	"fmt"
	"os"
	"task-tracker/internal/task/app"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		_, _ = fmt.Fprintln(os.Stderr, "need more arguments for todosher")
		os.Exit(1)
	}

	app.Run(args[1:])
}
