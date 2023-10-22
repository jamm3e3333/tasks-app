package handler

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jamm3e3333/tasks-app/internal/domain/entity"
	port "github.com/jamm3e3333/tasks-app/internal/domain/repository"
)

type GetDetailTaskHandler struct {
	tr port.TaskRepositoryPort
}

func NewGetDetailTaskHandler(tr port.TaskRepositoryPort) *GetDetailTaskHandler {
	return &GetDetailTaskHandler{
		tr,
	}
}

func (gdch *GetDetailTaskHandler) Execute(id string) (<-chan entity.TaskEntity, <-chan error) {
	taskId, err := uuid.Parse(id)

	if err != nil {
		ech := make(chan error)
		ch := make(chan entity.TaskEntity)

		go func() {
			defer close(ech)
			defer close(ch)
			e := errors.New(fmt.Sprintf("invalid task id: %s", id))
			ech <- e
			ch <- entity.TaskEntity{}
		}()
		return ch, ech
	}

	return gdch.tr.TaskById(taskId)
}
