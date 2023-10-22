package handler

import (
	"github.com/jamm3e3333/tasks-app/internal/application/command"
	"github.com/jamm3e3333/tasks-app/internal/domain/entity"
	port "github.com/jamm3e3333/tasks-app/internal/domain/repository"
)

type UpdateTaskHandler struct {
	tr port.TaskRepositoryPort
}

func NewUpdateTaskHandler(tr port.TaskRepositoryPort) *UpdateTaskHandler {
	return &UpdateTaskHandler{tr}
}

func (uth *UpdateTaskHandler) Execute(c command.UpdateTaskCommand) error {
	task := entity.NewTaskEntity(c.Task, c.IsDone, c.Id)

	return uth.tr.Update(*task)
}
