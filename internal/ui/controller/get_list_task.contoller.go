package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jamm3e3333/tasks-app/internal/application/query/handler"
	responseBody "github.com/jamm3e3333/tasks-app/internal/ui/response"
	"net/http"
)

type GetListTaskController struct {
	glch *handler.GetListTaskHandler
}

func NewGetListTask(glch *handler.GetListTaskHandler) *GetListTaskController {

	return &GetListTaskController{
		glch,
	}
}

func (glcc *GetListTaskController) Register(ctx *gin.Engine) {
	ctx.GET("/task", glcc.tasks)
}

// tasks fetches a list of tasks.
// @Summary Get a list of tasks
// @Description Fetches a list of all available tasks.
// @Tags Task
// @Accept  json
// @Produce  json
// @Param X-Request-UUID header string true "Request UUID"
// @Success 200 {object} []responseBody.GetListTaskResponse "Successfully fetched list of tasks"
// @Failure 404 {object} errors.HTTPError "Tasks not found"
// @Router /task [get]
func (glcc *GetListTaskController) tasks(ctx *gin.Context) {

	var response []responseBody.GetListTaskResponse

	ec, done := glcc.glch.Execute()

	for {
		select {
		case entity := <-ec:
			response = append(response, responseBody.New(entity))
		case <-done:
			if l := len(response); l == 0 {
				ctx.JSON(http.StatusOK, gin.H{"data": []string{}})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"data": response})
			return
		}
	}
}
