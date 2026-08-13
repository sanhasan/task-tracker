package task

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"task-tracker/internal/task/entity"
)

func (h *Handler) Update(args []string) (string, error) {
	if len(args) < 2 {
		return "", ErrFewArguments
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return "", errors.Join(ErrWrongArgument, err)
	}

	if strings.ToLower(args[1]) == "-cn" {
		return h.updateDescription(id, args[2:])
	}

	if len(args) > 2 {
		return "", ErrALotOFArguments
	}
	return h.updateStatus(id, args[1])
}

func (h *Handler) updateDescription(id int, args []string) (string, error) {
	if len(args) == 0 {
		return "", ErrFewArguments
	}

	newDesc := strings.Join(args, " ")
	if err := h.service.UpdateDescription(id, newDesc); err != nil {
		return "", fmt.Errorf("error while updating description task with id %d: %v", id, err)
	}

	resp := "update description task with id " + strconv.Itoa(id)
	return resp, nil
}

func (h *Handler) updateStatus(id int, status string) (string, error) {
	tStatus, err := entity.StringTOTaskStatus(status)

	if err != nil {
		return "", errors.Join(ErrWrongArgument, err)
	}

	if err = h.service.UpdateStatus(id, tStatus); err != nil {
		return "", fmt.Errorf("error while updating status into %s, task with id %d: %v", status, id, err)
	}
	resp := "update status into " + status + ", task with id " + strconv.Itoa(id)
	return resp, nil
}
