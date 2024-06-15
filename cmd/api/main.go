package main

import (
	"log"
	"onlyanotherblog/config"
	"onlyanotherblog/internal/app"

	"github.com/gin-gonic/gin"
)

func main() {
	serverConfig, err := config.LoadServerConfig()
	if err != nil {
		log.Fatal(err)
	}

	ginEngine := gin.Default()
	app := app.NewApp(ginEngine, &serverConfig)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
