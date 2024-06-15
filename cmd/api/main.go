package main

import (
	"log"
	"onlyanotherblog/config"
	database "onlyanotherblog/database/sqlc"
	"onlyanotherblog/internal/app"

	"github.com/gin-gonic/gin"
)

func main() {
	serverConfig, err := config.LoadServerConfig()
	if err != nil {
		log.Fatal(err)
	}

	databaseRepository, err := database.NewDatabaseRepository(*serverConfig)
	if err != nil {
		log.Fatal(err)
	}

	ginEngine := gin.Default()
	app := app.NewApp(ginEngine, serverConfig, databaseRepository)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
