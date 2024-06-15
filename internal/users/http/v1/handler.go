package v1

import (
	"net/http"
	database "onlyanotherblog/database/sqlc"
	"onlyanotherblog/internal/users/dtos"
	"onlyanotherblog/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type UsersHandler struct {
	DatabaseRepository *database.DatabaseRepository
}

func (uh *UsersHandler) createUser(c *gin.Context) {
	createUserBody := dtos.CreateUserRequestBody{}

	if err := c.BindJSON(&createUserBody); err != nil {
		responseError := utils.ResponseErrorStackTrace("occurred an erro in deserialization body data.", err.Error())
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	passwordHashed, err := utils.HashPassword(createUserBody.Password)
	if err != nil {
		responseError := utils.ResponseErrorMessage("occurred an error in create of user.")
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	argCreateUser := database.CreateUserParams{
		ID: utils.GenerateUniqueId(),
		Username: pgtype.Text{
			String: createUserBody.Username,
			Valid:  true,
		},
		Email: pgtype.Text{
			String: createUserBody.Email,
			Valid:  true,
		},
		Password: pgtype.Text{
			String: passwordHashed,
			Valid:  true,
		},
	}

	user, err := uh.DatabaseRepository.CreateUser(c, argCreateUser)
	if err != nil {
		responseError := utils.ResponseErrorMessage("occurred an error in create of user.")
		c.JSON(http.StatusInternalServerError, responseError)
		return
	}

	c.JSON(http.StatusOK, utils.ReponseData(user))
}
