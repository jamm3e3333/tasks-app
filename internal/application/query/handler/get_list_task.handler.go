package handler

import (
	"github.com/jamm3e3333/tasks-app/internal/domain/entity"
	port "github.com/jamm3e3333/tasks-app/internal/domain/repository"
)

type GetListTaskHandler struct {
	tr port.TaskRepositoryPort
}

func NewGetListTaskHandler(tr port.TaskRepositoryPort) *GetListTaskHandler {
	return &GetListTaskHandler{
		tr,
	}
}

func (glch *GetListTaskHandler) Execute() (<-chan entity.TaskEntity, <-chan bool) {
	return glch.tr.Tasks()
}
