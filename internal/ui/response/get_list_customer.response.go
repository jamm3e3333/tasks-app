package response

import (
	"github.com/google/uuid"
	"github.com/jamm3e3333/tasks-app/internal/domain/entity"
)

type GetListTaskResponse struct {
	Id     uuid.UUID `json:"id"`
	Task   string    `json:"task"`
	IsDone bool      `json:"is_done"`
}

func New(t entity.TaskEntity) GetListTaskResponse {
	return GetListTaskResponse{
		t.Id,
		t.Task,
		t.IsDone,
	}
}
