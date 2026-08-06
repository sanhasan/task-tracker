package useCase

import "fmt"

func (t *TaskService) Add(desc string) error {
	err := t.repo.AddEvent(desc)
	if err != nil {
		return fmt.Errorf("failed add event: %w", err)
	}
	return nil
}
