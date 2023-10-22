package internal

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jamm3e3333/tasks-app/docs" // I declare this unused import so swagger can run
	commandHandler "github.com/jamm3e3333/tasks-app/internal/application/command/handler"
	queryHandler "github.com/jamm3e3333/tasks-app/internal/application/query/handler"
	repository "github.com/jamm3e3333/tasks-app/internal/infra/in-memory"
	"github.com/jamm3e3333/tasks-app/internal/ui/controller"
	"github.com/jamm3e3333/tasks-app/pkg/logger"
	pkgGin "github.com/jamm3e3333/tasks-app/pkg/net/http/gin"
	ginprometheus "github.com/jamm3e3333/tasks-app/pkg/net/http/gin_prometheus"
	"github.com/jamm3e3333/tasks-app/pkg/net/http/health"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeApp(ge *gin.Engine, log logger.Logger) {
	ge.Any("/metrics", ginprometheus.Handler())
	ge.Use(pkgGin.LoggerMiddleware(pkgGin.NewLoggerMiddlewareConfig(
		"X-Request-UUID",
		[]string{"/metrics", "/health", "/status", "/docs/*any"},
	)))

	ge.GET("health", health.HandleHealthCheck)

	ge.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	tr := repository.NewTaskRepository(log)

	cth := commandHandler.NewCreateTaskHandler(tr)
	ctc := controller.NewCreateTaskController(cth)
	ctc.Register(ge)

	gdch := queryHandler.NewGetDetailTaskHandler(tr)
	gdc := controller.NewGetDetailTask(gdch)
	gdc.Register(ge)

	glch := queryHandler.NewGetListTaskHandler(tr)
	glc := controller.NewGetListTask(glch)
	glc.Register(ge)

	uth := commandHandler.NewUpdateTaskHandler(tr)
	utc := controller.NewUpdateTaskController(uth)
	utc.Register(ge)

	dth := commandHandler.NewDeleteTaskHandler(tr)
	dtc := controller.NewDeleteTaskController(dth)
	dtc.Register(ge)
}
