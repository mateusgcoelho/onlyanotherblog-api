package app

import (
	"fmt"
	"onlyanotherblog/config"
	v1Posts "onlyanotherblog/internal/posts/http/v1"
	v1Users "onlyanotherblog/internal/users/http/v1"

	database "onlyanotherblog/database/sqlc"

	"github.com/gin-contrib/cors"
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
	app.ginEngine.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Cross-Origin, Access-Control-Allow-Origin"},
	}))

	usersHandler := v1Users.UsersHandler{
		DatabaseRepository: app.databaseRepository,
	}
	postsHandler := v1Posts.PostsHandler{
		DatabaseRepository: app.databaseRepository,
	}

	userGroup := app.ginEngine.Group("/users")
	usersHandler.UserRoutes(userGroup)

	postGroup := app.ginEngine.Group("/posts")
	postsHandler.PostRoutes(postGroup)

	port := fmt.Sprintf(":%v", app.serverConfig.ServerPort)
	return app.ginEngine.Run(port)
}
