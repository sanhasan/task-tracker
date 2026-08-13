package task

import (
	"errors"
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
	id, err := strconv.Atoi(args[0])

	if err != nil {
		return "", errors.Join(ErrWrongArgument, err)
	}

	if err := h.service.Delete(id); err != nil {
		return "", fmt.Errorf("error while deleting task with id %d: %v", id, err)
	}

	resp := "delete task with id " + strconv.Itoa(id)
	return resp, nil
}
