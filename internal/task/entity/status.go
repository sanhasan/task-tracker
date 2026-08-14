package entity

import "fmt"

type TaskStatus uint8

const (
	StatusUndefined TaskStatus = iota
	StatusTodo
	StatusINProcess
	StatusDone
)

var allowedTransitions = map[TaskStatus][]TaskStatus{
	StatusTodo:      {StatusINProcess},
	StatusINProcess: {StatusINProcess, StatusDone},
	StatusDone:      {StatusINProcess},
}

func (t TaskStatus) CanTransitionTO(next TaskStatus) bool {
	for _, st := range allowedTransitions[t] {
		if st == next {
			return true
		}
	}
	return false
}

func (t TaskStatus) TOString() string {
	switch t {
	case StatusTodo:
		return "todo"
	case StatusINProcess:
		return "in-process"
	case StatusDone:
		return "done"
	default:
		return "undefined"
	}
}

func StringTOTaskStatus(status string) (TaskStatus, error) {
	switch status {
	case "todo":
		return StatusTodo, nil
	case "in-process":
		return StatusINProcess, nil
	case "done":
		return StatusDone, nil
	default:
		return StatusUndefined, fmt.Errorf("anavailable status")
	}
}
