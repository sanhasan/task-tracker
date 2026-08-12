package task

import (
	"fmt"
	"strconv"
)

func (h *Handler) Add(args []string) (string, error) {
	if len(args) == 0 {
		return "", fmt.Errorf("")
	}
	desc := args[0]

	id, err := h.service.Add(desc)
	if err != nil {
		return "", fmt.Errorf("error while adding a new task: %v", err)
	}

	resp := "add task, her id - " + strconv.Itoa(id)
	return resp, nil
}
