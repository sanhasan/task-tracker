package entity

import (
	"strconv"
	"strings"
	"time"
)

type Task struct {
	id          int
	description string
	status      TaskStatus
	createdAt   time.Time
	updatedAt   time.Time
}

func New(id int, description string) *Task {
	t := time.Now()
	return &Task{
		id:          id,
		description: description,
		status:      StatusTodo,
		createdAt:   t,
		updatedAt:   t,
	}
}

func (t *Task) ToString() string {
	task := make([]string, 5)
	task[0] = strconv.Itoa(t.id)
	task[1] = t.status.TOString()
	task[2] = t.description
	task[3] = t.updatedAt.String()
	task[4] = t.createdAt.String()

	return strings.Join(task, "   ")
}

func (t *Task) Status() TaskStatus {
	return t.status
}
