package v1

import (
	"onlyanotherblog/internal/auth/middlewares"

	"github.com/gin-gonic/gin"
)

func (ph *PostsHandler) PostRoutes(group *gin.Engine) {
	group.GET("/posts", ph.getPosts)
	group.GET("/posts/users/:username", ph.getPostsByUsername)
	group.GET("/posts/me", middlewares.MiddlewareAuth(ph.TokenMaker), ph.getPostsOfUser)
	group.GET("/posts/:id", ph.getPost)
	group.POST("/posts", middlewares.MiddlewareAuth(ph.TokenMaker), ph.createPost)
}
