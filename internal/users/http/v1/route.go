package v1

import "github.com/gin-gonic/gin"

func (uh *UsersHandler) UserRoutes(group *gin.RouterGroup) {
	group.POST("/", uh.createUser)
}
