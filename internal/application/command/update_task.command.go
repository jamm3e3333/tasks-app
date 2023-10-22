package command

import "github.com/google/uuid"

type UpdateTaskCommand struct {
	Id     uuid.UUID
	Task   string
	IsDone bool
}

func NewUpdateTaskCommand(taskId uuid.UUID, task string, isDone bool) UpdateTaskCommand {

	return UpdateTaskCommand{
		Id:     taskId,
		Task:   task,
		IsDone: isDone,
	}
}
