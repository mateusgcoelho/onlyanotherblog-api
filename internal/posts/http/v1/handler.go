package v1

import (
	"net/http"
	database "onlyanotherblog/database/sqlc"
	"onlyanotherblog/internal/posts/dtos"
	"onlyanotherblog/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type PostsHandler struct {
	DatabaseRepository *database.DatabaseRepository
}

func (ph *PostsHandler) getPosts(c *gin.Context) {
	getPostsQuery := dtos.GetPostsRequestQuery{}

	if err := c.BindQuery(&getPostsQuery); err != nil {
		responseError := utils.ResponseErrorStackTrace("occurred an erro in deserialization query params.", err.Error())
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	basedPostId := ""

	if getPostsQuery.BasedPostId != nil {
		basedPostId = *getPostsQuery.BasedPostId
	}

	argGetPosts := database.GetPostsParams{
		ID:    basedPostId,
		Limit: int32(getPostsQuery.ItensPerPage),
	}
	posts, err := ph.DatabaseRepository.GetPosts(c, argGetPosts)
	if err != nil {
		responseError := utils.ResponseErrorMessage("occurred an error in get posts.")
		c.JSON(http.StatusInternalServerError, responseError)
		return
	}

	c.JSON(http.StatusOK, utils.ReponseData(posts))
}

func (ph *PostsHandler) getPost(c *gin.Context) {
	getPostParam := dtos.GetPostRequestParam{}

	if err := c.BindUri(&getPostParam); err != nil {
		responseError := utils.ResponseErrorStackTrace("occurred an erro in deserialization id param.", err.Error())
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	post, err := ph.DatabaseRepository.GetPost(c, *getPostParam.PostId)
	if err != nil {
		responseError := utils.ResponseErrorMessage("occurred an error in get post.")
		c.JSON(http.StatusInternalServerError, responseError)
		return
	}

	c.JSON(http.StatusOK, utils.ReponseData(post))
}

func (ph *PostsHandler) createPost(c *gin.Context) {
	createPostBody := dtos.CreatePostRequestBody{}

	if err := c.BindJSON(&createPostBody); err != nil {
		responseError := utils.ResponseErrorStackTrace("occurred an erro in deserialization body data.", err.Error())
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	argCreatePost := database.CreatePostParams{
		ID: utils.GenerateUniqueId(),
		Title: pgtype.Text{
			String: createPostBody.Title,
			Valid:  true,
		},
		Content: pgtype.Text{
			String: createPostBody.Content,
			Valid:  true,
		},
		UserID: pgtype.Text{
			String: "01J0ERCQFR1Y3Q5X78JPEVPX4D", // Alterar para id logado posteriomente
			Valid:  true,
		},
	}

	post, err := ph.DatabaseRepository.CreatePost(c, argCreatePost)
	if err != nil {
		responseError := utils.ResponseErrorMessage("occurred an error in create of post.")
		c.JSON(http.StatusInternalServerError, responseError)
		return
	}

	c.JSON(http.StatusOK, utils.ReponseData(post))
}
