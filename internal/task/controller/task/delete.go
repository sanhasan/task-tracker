package task

import (
	"fmt"
	"strconv"
)

func (h *Handler) Delete(args []string) (string, error) {
	if len(args) == 0 {
		return "", ErrFewArguments
	}
	if len(args) > 1 {
		return "", ErrALotOFArguments
	}
	id, _ := strconv.Atoi(args[0])

	err := h.service.Delete(id)
	if err != nil {
		return "", fmt.Errorf("error while deleting task with id %d: %v", id, err)
	}

	resp := "delete task"
	return resp, nil
}
