package app

import (
	"fmt"
	"onlyanotherblog/config"

	"github.com/gin-gonic/gin"
)

type app struct {
	ginEngine    *gin.Engine
	serverConfig *config.ServerConfig
}

func NewApp(ginEngine *gin.Engine, serverConfig *config.ServerConfig) *app {
	return &app{
		ginEngine:    ginEngine,
		serverConfig: serverConfig,
	}
}

func (app *app) Run() error {
	port := fmt.Sprintf(":%v", app.serverConfig.ServerPort)
	return app.ginEngine.Run(port)
}
