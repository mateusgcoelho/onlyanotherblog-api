package v1

import (
	"onlyanotherblog/internal/auth/middlewares"

	"github.com/gin-gonic/gin"
)

func (uh *UsersHandler) UserRoutes(router *gin.Engine) {
	router.GET("/users/me", middlewares.MiddlewareAuth(uh.TokenMaker), uh.getUserByToken)
	router.POST("/users", uh.createUser)
}
