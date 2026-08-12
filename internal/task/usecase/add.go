package useCase

import "fmt"

func (t *taskService) Add(desc string) (int, error) {
	id, err := t.repo.AddEvent(desc)
	if err != nil {
		return 0, fmt.Errorf("failed add event: %w", err)
	}
	return id, nil
}
