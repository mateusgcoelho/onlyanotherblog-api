package dtos

type GetPostsOfUserParam struct {
	Username string `uri:"username" binding:"required"`
}
