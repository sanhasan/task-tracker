package useCase

import "task-tracker/internal/task/entity"

type Repo interface {
	AddEvent(description string) error
	DeleteEvent(id int) error
	UpdateStatus(id int) error
	PrintEventsList() ([]entity.Task, error)
}

type TaskService struct {
	repo Repo
}

func New(r Repo) *TaskService {
	return &TaskService{repo: r}
}
