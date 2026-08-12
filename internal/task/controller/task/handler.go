package task

import (
	"errors"

	"task-tracker/internal/task/entity"
)

var (
	ErrFewArguments    = errors.New("soo few arguments")
	ErrALotOFArguments = errors.New("a lot of arguments")
	ErrWrongArgument   = errors.New("wrong arguments")
)

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
