package v1

import "github.com/gin-gonic/gin"

func (uh *AuthHandler) AuthRoutes(router *gin.Engine) {
	router.POST("/auth", uh.signIn)
}
