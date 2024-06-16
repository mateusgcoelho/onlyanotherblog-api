package main

import (
	"log"
	"onlyanotherblog/config"
	database "onlyanotherblog/database/sqlc"
	"onlyanotherblog/internal/app"
	"onlyanotherblog/internal/auth/token"

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

	pasetoTokenMaker, err := token.NewPasetoMaker(serverConfig.SecretKeyToken)
	if err != nil {
		log.Fatal(err)
	}

	ginEngine := gin.Default()
	ginEngine.Use(CORSMiddleware())

	app := app.NewApp(ginEngine, serverConfig, databaseRepository, &pasetoTokenMaker)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
