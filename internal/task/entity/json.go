package entity

import (
	"encoding/json"
	"fmt"
	"time"
)

type jsonTask struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (t *Task) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonTask{
		ID:          t.id,
		Description: t.description,
		Status:      t.status.TOString(),
		CreatedAt:   t.createdAt,
		UpdatedAt:   t.updatedAt,
	})
}

func (t *Task) UnmarshalJSON(data []byte) error {
	var task jsonTask
	if err := json.Unmarshal(data, &task); err != nil {
		return fmt.Errorf("unmarshal task: %w", err)
	}

	status, err := StringTOTaskStatus(task.Status)
	if err != nil {
		return fmt.Errorf("parsing task status: %w", err)
	}

	t.id = task.ID
	t.description = task.Description
	t.status = status
	t.createdAt = task.CreatedAt
	t.updatedAt = task.UpdatedAt

	return nil
}
