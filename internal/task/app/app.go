package app

import (
	"errors"
	"fmt"
	"task-tracker/internal/task/config"
	"task-tracker/internal/task/controller"
	hand "task-tracker/internal/task/controller/task"
	repository "task-tracker/internal/task/repository/json"
	"task-tracker/internal/task/usecase"
)

func Run(cfg *config.Config, args []string) (string, error) {
	repo, err := repository.New(cfg.TasksFilePath)
	if err != nil {
		return "", fmt.Errorf("create repository: %v", err)
	}

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
			return "", fmt.Errorf("%v. Use '%s help' to see the supported commands", err, cfg.AppName)
		}
		return "", err
	}
	return resp, nil
}
