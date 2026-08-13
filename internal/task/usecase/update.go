package useCase

import (
	"fmt"
	"task-tracker/internal/task/entity"
)

func (t *taskService) UpdateDescription(id int, desc string) error {
	err := t.repo.UpdateTaskDescription(id, desc)
	if err != nil {
		return fmt.Errorf("failed update event description: %w", err)
	}
	return nil
}

func (t *taskService) UpdateStatus(id int, status entity.TaskStatus) error {
	task, err := t.repo.GetEvent(id)
	if err != nil {
		return fmt.Errorf("failed update event status: %w", err)
	}

	if !task.Status.CanTransitionTO(status) {
		return fmt.Errorf("failed switch current status %s into %s: %w",
			task.Status.ToString(), status.ToString(), err)
	}

	err = t.repo.UpdateTaskStatus(id, status)
	if err != nil {
		return fmt.Errorf("failed update event status: %w", err)
	}
	return nil
}
