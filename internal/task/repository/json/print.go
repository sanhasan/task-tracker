package repository

import "task-tracker/internal/task/entity"

func (r *repo) PrintEventsList() ([]entity.Task, error) {
	st, err := r.load()
	if err != nil {
		return nil, err
	}

	return st.Tasks, nil
}
