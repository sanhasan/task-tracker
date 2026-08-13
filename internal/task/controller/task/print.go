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
		if strings.ToLower(args[0]) == "-t" {
			if len(args) == 1 {
				return "", ErrFewArguments
			} else if len(args) > 2 {
				return "", ErrALotOFArguments
			} else {
				switch args[1] {
				case "todo":
					list, err = h.service.PrintTODO()
					name = "todo"
				case "in-process":
					list, err = h.service.PrintINProcess()
					name = "in process"
				case "done":
					list, err = h.service.PrintDone()
					name = "done"
				default:
					return "", ErrWrongArgument
				}
			}
		} else {
			return "", ErrWrongArgument
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
