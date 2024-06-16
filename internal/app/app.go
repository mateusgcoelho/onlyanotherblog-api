package app

import (
	"fmt"
	"onlyanotherblog/config"
	v1Auth "onlyanotherblog/internal/auth/http/v1"
	"onlyanotherblog/internal/auth/token"
	v1Posts "onlyanotherblog/internal/posts/http/v1"
	v1Users "onlyanotherblog/internal/users/http/v1"

	database "onlyanotherblog/database/sqlc"

	"github.com/gin-gonic/gin"
)

type app struct {
	ginEngine          *gin.Engine
	serverConfig       *config.ServerConfig
	databaseRepository *database.DatabaseRepository
	tokenMaker         *token.Maker
}

func NewApp(
	ginEngine *gin.Engine,
	serverConfig *config.ServerConfig,
	databaseRepository *database.DatabaseRepository,
	tokenMaker *token.Maker,
) *app {
	return &app{
		ginEngine:          ginEngine,
		serverConfig:       serverConfig,
		databaseRepository: databaseRepository,
		tokenMaker:         tokenMaker,
	}
}

func (app *app) Run() error {
	authHandler := v1Auth.AuthHandler{
		DatabaseRepository: app.databaseRepository,
		TokenMaker:         *app.tokenMaker,
	}
	usersHandler := v1Users.UsersHandler{
		DatabaseRepository: app.databaseRepository,
		TokenMaker:         *app.tokenMaker,
	}
	postsHandler := v1Posts.PostsHandler{
		DatabaseRepository: app.databaseRepository,
	}

	authHandler.AuthRoutes(app.ginEngine)
	usersHandler.UserRoutes(app.ginEngine)
	postsHandler.PostRoutes(app.ginEngine)

	port := fmt.Sprintf(":%v", app.serverConfig.ServerPort)
	return app.ginEngine.Run(port)
}
