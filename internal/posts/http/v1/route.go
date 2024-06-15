package v1

import (
	"github.com/gin-gonic/gin"
)

func (ph *PostsHandler) PostRoutes(group *gin.RouterGroup) {
	group.GET("/", ph.getPosts)
	group.GET("/:id", ph.getPost)
	group.POST("/", ph.createPost)
}
