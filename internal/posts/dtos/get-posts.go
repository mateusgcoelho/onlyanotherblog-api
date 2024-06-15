package dtos

type GetPostsRequestQuery struct {
	BasedPostId  *string `form:"based_post_id"`
	ItensPerPage int     `form:"itens_per_page" binding:"required"`
}

type GetPostRequestParam struct {
	PostId *string `uri:"id"`
}
