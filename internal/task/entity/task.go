package entity

import "time"

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
