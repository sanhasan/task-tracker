package controller

import "errors"

var (
	ErrUnsupportedOperation = errors.New("unsupported operation")
	ErrZeroArguments        = errors.New("choose operation")
)

func (h *api) ParseCommand(args []string) (string, error) {
	if len(args) == 0 {
		return "", ErrZeroArguments
	}
	switch args[0] {
	case "help":
		return h.handler.Help(args[1:])
	case "add":
		return h.handler.Add(args[1:])
	case "delete":
		return h.handler.Delete(args[1:])
	case "update":
		return h.handler.Update(args[1:])
	case "print":
		return h.handler.Print(args[1:])
	}
	return "", ErrUnsupportedOperation
}
