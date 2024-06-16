package v1

import (
	"net/http"
	database "onlyanotherblog/database/sqlc"
	"onlyanotherblog/internal/auth/dtos"
	"onlyanotherblog/internal/auth/token"
	"onlyanotherblog/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type AuthHandler struct {
	DatabaseRepository *database.DatabaseRepository
	TokenMaker         token.Maker
}

func (uh *AuthHandler) signIn(c *gin.Context) {
	signInUserBody := dtos.SignInUserRequestBody{}

	if err := c.BindJSON(&signInUserBody); err != nil {
		responseError := utils.ResponseErrorStackTrace("occurred an erro in deserialization body data.", err.Error())
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	argGetUserByEmail := pgtype.Text{
		String: signInUserBody.Email,
		Valid:  true,
	}
	user, err := uh.DatabaseRepository.GetUserByEmail(c, argGetUserByEmail)
	if err != nil {
		responseError := utils.ResponseErrorMessage("email or password invalid, try again.")
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	isValidPassword := utils.ComparePasswords(user.Password.String, signInUserBody.Password)
	if !isValidPassword {
		responseError := utils.ResponseErrorMessage("email or password invalid, try again.")
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	token, err := uh.TokenMaker.CreateToken(user.ID, 48*time.Hour)
	if err != nil {
		responseError := utils.ResponseErrorMessage("occurred an error in auth.")
		c.JSON(http.StatusInternalServerError, responseError)
		return
	}

	response := dtos.SignInResponse{
		User: dtos.UserResponse{
			Id:        user.ID,
			Username:  user.Username.String,
			Email:     user.Email.String,
			CreatedAt: user.CreatedAt.Time,
			UpdatedAt: user.UpdatedAt.Time,
		},
		Token: token,
	}

	c.JSON(http.StatusOK, utils.ReponseData(response))
}
