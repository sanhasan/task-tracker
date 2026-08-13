package repository

import "task-tracker/internal/task/entity"

type repo struct {
}

func New() *repo {
	return &repo{}
}

func (r *repo) AddEvent(description string) (int, error) {
	return 0, nil
}

func (r *repo) DeleteEvent(id int) error {
	return nil
}

func (r *repo) GetEvent(id int) (entity.Task, error) {
	return *entity.New(id, ""), nil
}

func (r *repo) UpdateTaskDescription(id int, desc string) error {
	return nil
}
func (r *repo) UpdateTaskStatus(id int, status entity.TaskStatus) error {
	return nil
}

func (r *repo) PrintEventsList() ([]entity.Task, error) {
	return nil, nil
}
