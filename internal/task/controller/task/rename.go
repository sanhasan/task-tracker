package task

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func (h *Handler) Rename(args []string) (string, error) {
	if len(args) < 2 {
		return "", ErrFewArguments
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return "", errors.Join(ErrWrongArgument, err)
	}

	newDesc := strings.Join(args[1:], " ")
	if err := h.service.UpdateDescription(id, newDesc); err != nil {
		return "", fmt.Errorf("error while updating description task with id %d: %v", id, err)
	}

	resp := "rewrite description task with id " + strconv.Itoa(id)
	return resp, nil
}
