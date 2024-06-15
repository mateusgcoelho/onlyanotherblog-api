package app

import (
	"fmt"
	"onlyanotherblog/config"
	database "onlyanotherblog/database/sqlc"
	v1 "onlyanotherblog/internal/users/http/v1"

	"github.com/gin-gonic/gin"
)

type app struct {
	ginEngine          *gin.Engine
	serverConfig       *config.ServerConfig
	databaseRepository *database.DatabaseRepository
}

func NewApp(ginEngine *gin.Engine, serverConfig *config.ServerConfig, databaseRepository *database.DatabaseRepository) *app {
	return &app{
		ginEngine:          ginEngine,
		serverConfig:       serverConfig,
		databaseRepository: databaseRepository,
	}
}

func (app *app) Run() error {
	usersHandler := v1.UsersHandler{
		DatabaseRepository: app.databaseRepository,
	}

	userGroup := app.ginEngine.Group("/users")
	usersHandler.UserRoutes(userGroup)

	port := fmt.Sprintf(":%v", app.serverConfig.ServerPort)
	return app.ginEngine.Run(port)
}
