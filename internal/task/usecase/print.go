package useCase

import (
	"fmt"
	"task-tracker/internal/task/entity"
)

func (t *taskService) PrintAll() ([]entity.Task, error) {
	tasks, err := t.repo.PrintEventsList()

	if err != nil {
		return nil, fmt.Errorf("failed read all tasks: %w", err)
	}
	return tasks, nil
}

func (t *taskService) PrintToDo() ([]entity.Task, error) {
	status := "todo"
	return t.print(func(st entity.TaskStatus) bool {
		if st.ToString() == status {
			return true
		}
		return false
	}, status)
}

func (t *taskService) PrintINProcess() ([]entity.Task, error) {
	status := "in process"
	return t.print(func(st entity.TaskStatus) bool {
		if st.ToString() == status {
			return true
		}
		return false
	}, status)
}

func (t *taskService) PrintDone() ([]entity.Task, error) {
	status := "done"
	return t.print(func(st entity.TaskStatus) bool {
		if st.ToString() == status {
			return true
		}
		return false
	}, status)
}

func (t *taskService) print(filter func(task entity.TaskStatus) bool, name string) ([]entity.Task, error) {
	tasks, err := t.repo.PrintEventsList()

	if err != nil {
		return nil, fmt.Errorf("failed read %s tasks: %w", name, err)
	}

	res := make([]entity.Task, 0)

	for _, v := range tasks {
		if filter(v.Status) {
			res = append(res, v)
		}
	}
	return res, nil
}
