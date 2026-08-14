package task

import (
	"errors"
	"fmt"
	"strconv"
	"task-tracker/internal/task/entity"
)

func (h *Handler) Update(args []string) (string, error) {
	if len(args) < 2 {
		return "", ErrFewArguments
	}
	if len(args) > 2 {
		return "", ErrALotOFArguments
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return "", errors.Join(ErrWrongArgument, err)
	}

	return h.updateStatus(id, args[1])
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
