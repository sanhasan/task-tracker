package repository

import "task-tracker/internal/task/entity"

func (r *repo) AddEvent(description string) (int, error) {
	st, err := r.load()
	if err != nil {
		return 0, err
	}

	taskId := st.NextID
	task := entity.New(taskId, description)

	st.NextID++
	st.Tasks = append(st.Tasks, *task)

	if err := r.save(st); err != nil {
		return 0, err
	}
	return taskId, nil
}
