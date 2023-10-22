package handler

import (
	"github.com/google/uuid"
	"github.com/jamm3e3333/tasks-app/internal/application/command"
	"github.com/jamm3e3333/tasks-app/internal/domain/entity"
	port "github.com/jamm3e3333/tasks-app/internal/domain/repository"
)

type CreateTaskHandler struct {
	tr port.TaskRepositoryPort
}

func NewCreateTaskHandler(tr port.TaskRepositoryPort) *CreateTaskHandler {
	return &CreateTaskHandler{
		tr,
	}
}

func (cth *CreateTaskHandler) Execute(ctc command.CreateTaskCommand) error {
	id := uuid.New()
	task := entity.NewTaskEntity(ctc.Task, ctc.IsDone, id)

	return cth.tr.Create(*task)
}
