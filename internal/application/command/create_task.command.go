package command

type CreateTaskCommand struct {
	Task   string
	IsDone bool
}

func NewCreateTaskCommand(task string, isDone bool) CreateTaskCommand {
	return CreateTaskCommand{
		task,
		isDone,
	}
}
