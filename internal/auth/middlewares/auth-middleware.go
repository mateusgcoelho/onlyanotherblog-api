package middlewares

import (
	"net/http"
	"onlyanotherblog/internal/auth/token"
	"onlyanotherblog/pkg/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func MiddlewareAuth(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerToken := ctx.GetHeader("Authorization")

		value := strings.Split(headerToken, " ")
		if len(value) < 2 || value[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, utils.ResponseErrorMessage("bad token format."))
			return
		}

		payload, err := tokenMaker.VerifyToken(value[1])
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, utils.ResponseErrorMessage(err.Error()))
			return
		}

		ctx.Set("user_id", payload.UserId)
		ctx.Next()
	}
}
