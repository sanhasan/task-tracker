package repository

import (
	"task-tracker/internal/task/entity"
)

func (r *repo) GetEvent(id int) (entity.Task, error) {
	st, err := r.load()
	if err != nil {
		return entity.Task{}, err
	}

	pos := st.findIndex(id)
	if pos == -1 {
		return entity.Task{}, ErrEventNotFound
	}
	return st.Tasks[pos], nil
}
