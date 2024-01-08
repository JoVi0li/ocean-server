package auth

import (
	"net/http"

	"github.com/JoVi0li/ocean-server/internal/util"
	"github.com/gin-gonic/gin"
)

func AuthMidd() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := util.GetToken(ctx)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"sucess": false,
				"data":   nil,
				"error":  err,
			})
			ctx.Abort()

			return
		}

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
