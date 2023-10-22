package body

type CreateTaskBody struct {
	Task   string `json:"task" binding:"required"`
	IsDone bool   `json:"is_done" binding:"required"`
}
