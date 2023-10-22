package entity

import (
	"github.com/google/uuid"
)

type TaskEntity struct {
	Id     uuid.UUID
	Task   string
	IsDone bool
}

func NewTaskEntity(taskTodo string, isDone bool, id uuid.UUID) *TaskEntity {
	return &TaskEntity{
		Id:     id,
		Task:   taskTodo,
		IsDone: isDone,
	}
}
