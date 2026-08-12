package useCase

import "fmt"

func (t *taskService) Delete(id int) error {
	err := t.repo.DeleteEvent(id)
	if err != nil {
		return fmt.Errorf("failed delete event: %w", err)
	}
	return nil
}
