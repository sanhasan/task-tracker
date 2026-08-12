package entity

import (
	"strconv"
	"strings"
	"time"
)

type TaskStatus uint8

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
	case 0:
		return "todo"
	case 1:
		return "in process"
	case 2:
		return "done"
	default:
		return "undefined"
	}
}
