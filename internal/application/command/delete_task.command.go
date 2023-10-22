package command

import "github.com/google/uuid"

type DeleteTaskCommand struct {
	Id uuid.UUID
}

func NewDeleteTaskCommand(taskId uuid.UUID) DeleteTaskCommand {

	return DeleteTaskCommand{
		Id: taskId,
	}
}
