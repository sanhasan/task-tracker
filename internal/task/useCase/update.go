package useCase

import "fmt"

func (t *TaskService) Update(id int) error {
	err := t.repo.UpdateStatus(id)
	if err != nil {
		return fmt.Errorf("failed update event status: %w", err)
	}
	return nil
}
