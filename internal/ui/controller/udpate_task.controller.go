package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jamm3e3333/tasks-app/internal/application/command"
	"github.com/jamm3e3333/tasks-app/internal/application/command/handler"
	reqBody "github.com/jamm3e3333/tasks-app/internal/ui/body"
	"github.com/jamm3e3333/tasks-app/pkg/net/http/errors"
	"net/http"
)

type UpdateTaskController struct {
	uth *handler.UpdateTaskHandler
}

func NewUpdateTaskController(uth *handler.UpdateTaskHandler) *UpdateTaskController {
	return &UpdateTaskController{uth}
}

func (utc *UpdateTaskController) Register(ge *gin.Engine) {
	ge.PUT("/task/:id", utc.update)
}

// update updates a task.
// @Summary Update a task
// @Description Updates the details of a task identified by its ID.
// @Tags Task
// @Accept  json
// @Produce  json
// @Param id path string true "Task ID"
// @Param X-Request-UUID header string true "Request UUID"
// @Param body body reqBody.UpdateTaskBody true "Task update payload"
// @Success 200 {object} map[string]interface{} "Successfully updated task"
// @Failure 400 {object} errors.HTTPError "Bad request"
// @Router /task/{id} [put]
func (utc *UpdateTaskController) update(ctx *gin.Context) {
	var body reqBody.UpdateTaskBody
	var taskUuid uuid.UUID

	taskId, _ := ctx.Params.Get("id")
	err := ctx.BindJSON(&body)

	badRequest := errors.NewBadRequest(fmt.Sprintf("cannot update task with id %s", taskId))

	if err != nil {
		ctx.JSON(badRequest.HTTPCode(), badRequest.JSON())
		return
	}

	taskUuid, err = uuid.Parse(taskId)
	if err != nil {
		ctx.JSON(badRequest.HTTPCode(), badRequest.JSON())
		return
	}

	cmd := command.NewUpdateTaskCommand(taskUuid, body.Task, body.IsDone)
	err = utc.uth.Execute(cmd)
	if err != nil {
		ctx.JSON(badRequest.HTTPCode(), badRequest.JSON())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "task updated"})
}
