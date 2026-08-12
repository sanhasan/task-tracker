package repository

import "task-tracker/internal/task/entity"

type repo struct {
}

func New() *repo {
	return &repo{}
}

func (r *repo) AddEvent(description string) error {

}

func (r *repo) DeleteEvent(id int) error {

}

func (r *repo) UpdateStatus(id int) error {

}

func (r *repo) PrintEventsList() ([]entity.Task, error) {

}
