package repository

import "task-tracker/internal/task/entity"

func (r *repo) UpdateTaskStatus(id int, status entity.TaskStatus) error {
	st, err := r.load()
	if err != nil {
		return err
	}

	pos := st.findIndex(id)
	if pos == -1 {
		return ErrEventNotFound
	}

	st.Tasks[pos].UpdateStatus(status)
	r.swapTask(pos, st)
	return r.save(st)
}
