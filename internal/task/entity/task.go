package entity

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type TaskStatus uint8

const (
	StatusUndefined TaskStatus = iota
	StatusTodo
	StatusINProcess
	StatusDone
)

type Task struct {
	id          int
	description string
	Status      TaskStatus
	createdAt   time.Time
	updatedAt   time.Time
}

func New(id int, description string) *Task {
	t := time.Now()
	return &Task{
		id:          id,
		description: description,
		Status:      0,
		createdAt:   t,
		updatedAt:   t,
	}
}

func (t *Task) ToString() string {
	task := make([]string, 5)
	task[0] = strconv.Itoa(t.id)
	task[1] = t.Status.ToString()
	task[2] = t.description
	task[3] = t.updatedAt.String()
	task[4] = t.createdAt.String()

	return strings.Join(task, "   ")
}

func (t TaskStatus) ToString() string {
	switch t {
	case StatusTodo:
		return "todo"
	case StatusINProcess:
		return "in process"
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
	case "in process":
		return StatusINProcess, nil
	case "done":
		return StatusDone, nil
	default:
		return StatusUndefined, fmt.Errorf("anavailable status")
	}
}

var allowedTransitions = map[TaskStatus][]TaskStatus{
	StatusTodo:      {StatusINProcess},
	StatusINProcess: {StatusINProcess, StatusDone},
	StatusDone:      {StatusINProcess},
}

func (s TaskStatus) CanTransitionTO(next TaskStatus) bool {
	for _, st := range allowedTransitions[s] {
		if st == next {
			return true
		}
	}
	return false
}
