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
			String: "01J0E7W3BV9M7YQ3YJB0VYFSS8",
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
