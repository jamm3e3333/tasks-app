package health

import (
	httppkg "net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Health check
// @Description Health check of the application
// @Tags App
// @Produce json
// @Success 204
// @Router /health [get]
func HandleHealthCheck(context *gin.Context) {
	// TODO: add DB check
	context.Status(httppkg.StatusNoContent)
}
