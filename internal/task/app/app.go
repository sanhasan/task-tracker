package app

import (
	"errors"
	"fmt"
	"task-tracker/internal/task/controller"
	hand "task-tracker/internal/task/controller/task"
	repository "task-tracker/internal/task/repository/json"
	"task-tracker/internal/task/usecase"
)

func Run(args []string) (string, error) {

	repo := repository.New()
	service := useCase.New(repo)
	handler := hand.New(service)
	app := controller.New(handler)

	resp, err := app.ParseCommand(args)

	if err != nil {
		if errors.Is(err, controller.ErrUnsupportedOperation) ||
			errors.Is(err, controller.ErrZeroArguments) ||
			errors.Is(err, hand.ErrFewArguments) ||
			errors.Is(err, hand.ErrALotOFArguments) ||
			errors.Is(err, hand.ErrWrongArgument) {
			return "", fmt.Errorf("%v. Use todosher help to see the supported commands", err)
		}
		return "", err
	}
	return resp, nil
}
