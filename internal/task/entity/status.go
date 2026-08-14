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

func StringTOTaskStatus(status string) (TaskStatus, error) {
	switch status {
	case "todo":
		return StatusTodo, nil
	case "in process":
		return StatusINProcess, nil
	case "done":
		return StatusDone, nil
	default:
		return StatusUndefined, fmt.Errorf("anavailable status")
	}
}
