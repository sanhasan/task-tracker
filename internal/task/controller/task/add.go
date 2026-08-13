package task

import (
	"fmt"
	"strconv"
	"strings"
)

func (h *Handler) Add(args []string) (string, error) {
	if len(args) == 0 {
		return "", ErrFewArguments
	}

	desc := strings.Join(args, " ")

	id, err := h.service.Add(desc)
	if err != nil {
		return "", fmt.Errorf("error while adding a new task: %v", err)
	}

	resp := "add task, her id - " + strconv.Itoa(id)
	return resp, nil
}
