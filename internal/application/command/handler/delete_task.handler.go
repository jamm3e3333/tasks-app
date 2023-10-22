package handler

import (
	"github.com/jamm3e3333/tasks-app/internal/application/command"
	"github.com/jamm3e3333/tasks-app/internal/domain/repository"
)

type DeleteTaskHandler struct {
	tr repository.TaskRepositoryPort
}

func NewDeleteTaskHandler(tr repository.TaskRepositoryPort) *DeleteTaskHandler {
	return &DeleteTaskHandler{
		tr,
	}
}

func (dth *DeleteTaskHandler) Execute(command command.DeleteTaskCommand) error {
	if err := dth.tr.DeleteById(command.Id); err != nil {
		return err
	}
	return nil
}
