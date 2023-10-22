package vo

import "github.com/google/uuid"

type TaskUuidVo struct {
	Id uuid.UUID `json:"id"`
}

func (id *TaskUuidVo) New() *TaskUuidVo {
	return &TaskUuidVo{
		Id: uuid.New(),
	}
}

func (id *TaskUuidVo) TaskId() uuid.UUID {
	return id.Id
}
