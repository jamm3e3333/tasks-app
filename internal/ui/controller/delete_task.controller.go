package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jamm3e3333/tasks-app/internal/application/command"
	"github.com/jamm3e3333/tasks-app/internal/application/command/handler"
	"github.com/jamm3e3333/tasks-app/pkg/net/http/errors"
	"net/http"
)

type DeleteTaskController struct {
	dth *handler.DeleteTaskHandler
}

func NewDeleteTaskController(dth *handler.DeleteTaskHandler) *DeleteTaskController {
	return &DeleteTaskController{
		dth,
	}
}

func (dtc *DeleteTaskController) Register(ge *gin.Engine) {
	ge.DELETE("/task/:id", dtc.delete)
}

// delete deletes a task by its ID.
// @Summary Delete a task
// @Description Deletes a task identified by its ID.
// @Tags Task
// @Accept  json
// @Produce  json
// @Param id path string true "Task ID"
// @Param X-Request-UUID header string true "Request UUID"
// @Success 200 {object} errors.HTTPError "Successfully deleted task"
// @Failure 400 {object} errors.HTTPError "Bad request - Cannot delete task"
// @Failure 500 {object} errors.HTTPError "Internal server error"
// @Router /task/{id} [delete]
func (dtc *DeleteTaskController) delete(ctx *gin.Context) {
	taskId, _ := ctx.Params.Get("id")
	internalError := errors.NewInternalServerError("internal")

	taskUuid, err := uuid.Parse(taskId)
	if err != nil {
		ctx.JSON(internalError.HTTPCode(), internalError.JSON())
		return
	}

	deleteCommand := command.NewDeleteTaskCommand(taskUuid)
	err = dtc.dth.Execute(deleteCommand)
	if err != nil {
		badRequest := errors.NewBadRequest(err.Error())
		ctx.JSON(badRequest.HTTPCode(), badRequest.JSON())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "task deleted"})
}
