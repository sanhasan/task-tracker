package useCase

import "fmt"

func (t *taskService) Add(desc string) error {
	err := t.repo.AddEvent(desc)
	if err != nil {
		return fmt.Errorf("failed add event: %w", err)
	}
	return nil
}
