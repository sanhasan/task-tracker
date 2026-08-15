package task

import "errors"

const helpText = `todosher - simple command-line task tracker

USAGE:
  todosher <command> [arguments]

COMMANDS:
  add <description>                 add a new task
  delete id                         delete a task by its id
  update id status                  update task status
  rename id <description>           rewrite description task by its id
  print                             show all tasks
  print status                      show tasks with the given status
  help                              show this guide

STATUSES:
  todo                              tasks that are not started yet
  in-process                        tasks currently in progress
  done                              completed tasks

EXAMPLES:
  todosher add "buy milk"
  todosher add buy milk and bread
  todosher delete 2
  todosher print
  todosher print -t done

NOTE:
  description can be written with or without quotes:
  todosher add "buy milk"  is the same as  todosher add buy milk
`

func (h *Handler) Help(args []string) (string, error) {
	if len(args) > 0 {
		return "", errors.New("help command does not accept arguments")
	}

	return helpText, nil
}
