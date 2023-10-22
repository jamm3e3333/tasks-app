package main

import (
	"github.com/gin-gonic/gin"
	appConfig "github.com/jamm3e3333/tasks-app/config"
	"github.com/jamm3e3333/tasks-app/internal"
	"github.com/jamm3e3333/tasks-app/pkg/logger"
)

func main() {
	config := appConfig.CreateConfig()
	log := logger.New(config.LogLevel(), config.DevelMode())

	log.Info("config", config)
	server := gin.New()

	internal.InitializeApp(server, log)

	err := server.Run(":3000")

	if err != nil {
		panic("server not running")
	}
}
