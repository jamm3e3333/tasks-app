package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jamm3e3333/tasks-app/internal/application/query/handler"
	"github.com/jamm3e3333/tasks-app/pkg/net/http/errors"
	"net/http"
)

type GetDetailTaskController struct {
	gdch *handler.GetDetailTaskHandler
}

func NewGetDetailTask(gdch *handler.GetDetailTaskHandler) *GetDetailTaskController {
	return &GetDetailTaskController{
		gdch,
	}
}

func (gdcc *GetDetailTaskController) Register(ge *gin.Engine) {
	ge.GET("/task/:id", gdcc.detail)
}

// detail fetches the details of a specific task by ID.
// @Summary Get details of a specific task
// @Description Fetches detailed information for a specific task by its ID.
// @Tags Task
// @Accept  json
// @Produce  json
// @Param id path string true "Task ID"
// @Param X-Request-UUID header string true "Request UUID"
// @Success 200 {object} errors.HTTPError "Successfully fetched task details"
// @Failure 404 {object} errors.HTTPError "Task not found"
// @Router /task/{id} [get]
func (gdcc *GetDetailTaskController) detail(ctx *gin.Context) {
	taskId, _ := ctx.Params.Get("id")
	ch, ech := gdcc.gdch.Execute(taskId)

	select {
	case err := <-ech:
		notFound := errors.NewNotFound(err.Error())
		ctx.JSON(notFound.HTTPCode(), notFound.JSON())

	case task := <-ch:
		ctx.JSON(http.StatusOK, gin.H{"data": task})
	}
}
