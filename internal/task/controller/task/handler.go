package task

import "task-tracker/internal/task/entity"

type Service interface {
	Add(string) (int, error)
	Delete(int) error
	Update(int) error
	PrintAll() ([]entity.Task, error)
	PrintToDo() ([]entity.Task, error)
	PrintINProcess() ([]entity.Task, error)
	PrintDone() ([]entity.Task, error)
}

type Handler struct {
	service Service
}

func New(service Service) *Handler {
	return &Handler{
		service: service,
	}
}
