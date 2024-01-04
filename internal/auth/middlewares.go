package auth

import (
	"net/http"
	"strings"

	"github.com/JoVi0li/ocean-server/internal/util"
	"github.com/gin-gonic/gin"
)

func AuthMidd() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"sucess": false,
				"data":   nil,
				"error":  ErrorMissingAuthorizationToken,
			})
			ctx.Abort()

			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		tokenErr := util.ValidateToken(tokenString)

		if tokenErr != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"sucess": false,
				"data":   nil,
				"error":  tokenErr,
			})
			ctx.Abort()

			return
		}

		ctx.Next()
	}
}
