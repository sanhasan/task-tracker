package controller

import (
	handler "task-tracker/internal/task/controller/task"
)

type api struct {
	handler *handler.Handler
}

func New(service *handler.Handler) *api {
	return &api{
		handler: service,
	}
}
