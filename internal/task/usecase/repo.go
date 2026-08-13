package useCase

import "task-tracker/internal/task/entity"

type Repo interface {
	AddEvent(description string) (int, error)
	DeleteEvent(id int) error
	GetEvent(id int) (entity.Task, error)
	UpdateTaskDescription(id int, desc string) error
	UpdateTaskStatus(id int, status entity.TaskStatus) error
	PrintEventsList() ([]entity.Task, error)
}

type taskService struct {
	repo Repo
}

func New(r Repo) *taskService {
	return &taskService{repo: r}
}
