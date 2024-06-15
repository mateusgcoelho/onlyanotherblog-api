package v1

import "github.com/gin-gonic/gin"

func (ph *PostsHandler) PostRoutes(group *gin.RouterGroup) {
	group.POST("/", ph.createPost)
}
