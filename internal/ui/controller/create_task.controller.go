package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jamm3e3333/tasks-app/internal/application/command"
	"github.com/jamm3e3333/tasks-app/internal/application/command/handler"
	"github.com/jamm3e3333/tasks-app/internal/ui/body"
	"github.com/jamm3e3333/tasks-app/pkg/net/http/errors"
	"net/http"
)

type CreateTaskController struct {
	cth *handler.CreateTaskHandler
}

func NewCreateTaskController(cth *handler.CreateTaskHandler) *CreateTaskController {
	return &CreateTaskController{cth}
}

func (ctc *CreateTaskController) Register(ge *gin.Engine) {
	ge.POST("/task", ctc.create)
}

// create handles the task creation endpoint.
// @Summary Create a task
// @Description Create a new task with given details.
// @Tags Task
// @Accept  json
// @Produce  json
// @Param task body body.CreateTaskBody true "Task creation payload"
// @Param X-Request-UUID header string true "Request UUID"
// @Success 201 {object} map[string]string "Successfully created task"
// @Failure 400 {object} errors.HTTPError "Bad request"
// @Failure 500 {object} errors.HTTPError "Internal server error"
// @Router /task [post]
func (ctc *CreateTaskController) create(ctx *gin.Context) {
	var task body.CreateTaskBody
	badRequest := errors.NewBadRequest(fmt.Sprintf("cannot create task"))
	internalServerError := errors.NewInternalServerError("internal server error for creating task")

	if err := ctx.BindJSON(&task); err != nil {
		ctx.JSON(badRequest.HTTPCode(), badRequest.JSON())
		return
	}

	err := ctc.cth.Execute(command.NewCreateTaskCommand(task.Task, task.IsDone))
	if err != nil {
		ctx.JSON(internalServerError.HTTPCode(), internalServerError.JSON())
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "task created"})
}
