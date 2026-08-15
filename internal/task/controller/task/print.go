package task

import (
	"fmt"
	"strings"
	"task-tracker/internal/task/entity"
)

func (h *Handler) Print(args []string) (string, error) {
	var (
		list []entity.Task
		err  error
		name string
	)

	if len(args) == 0 {
		list, err = h.service.PrintAll()
	} else {
		if len(args) > 1 {
			return "", ErrALotOFArguments
		} else {
			name = args[0]
			switch name {
			case "todo":
				list, err = h.service.PrintTODO()
			case "in-process":
				list, err = h.service.PrintINProcess()
			case "done":
				list, err = h.service.PrintDone()
			default:
				return "", ErrWrongArgument
			}
		}
	}

	if err != nil {
		return "", fmt.Errorf("error while printing %stasks list: %v", name, err)
	}

	sb := strings.Builder{}
	for _, v := range list {
		sb.WriteString(v.ToString())
		sb.WriteString("\n")
	}
	resp := sb.String()
	return resp, nil
}
