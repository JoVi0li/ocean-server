package user

import (
	"context"
	"net/http"
	"time"

	"github.com/JoVi0li/ocean-server/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var service Service

func Configure() {
	service = Service{
		Repository: &RepositoryPostgres{
			Connection: database.Conn,
		},
	}
}

func GetUsers(ctx *gin.Context) {
	param := ctx.Param("id")

	if param == "" {
		ctx.JSON(http.StatusNotFound, gin.H{
			"sucess": false,
			"data":   nil,
			"error":  ErrorIdInvalid.Error(),
		})

		return
	}

	parsedId, err := uuid.Parse(param)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"sucess": false,
			"data":   nil,
			"error":  ErrorIdInvalid.Error(),
		})

		return
	}

	user, err := service.FindById(ctx, parsedId)

	if err != nil {
		statusCode := http.StatusInternalServerError

		if err == ErrorUserNotFound {
			statusCode = http.StatusNotFound
		}

		ctx.JSON(statusCode, gin.H{
			"sucess": false,
			"data":   nil,
			"error":  ErrorIdInvalid.Error(),
		})

		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"sucess": true,
		"data":   user,
		"error":  nil,
	})
}

func DeleteUsers(ctx *gin.Context) {
	param := ctx.Param("id")

	if param == "" {
		ctx.JSON(http.StatusNotFound, gin.H{
			"sucess": false,
			"data":   nil,
			"error":  ErrorIdInvalid.Error(),
		})

		return
	}

	parsedId, err := uuid.Parse(param)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"sucess": false,
			"data":   nil,
			"error":  ErrorIdInvalid.Error(),
		})

		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := service.DeleteById(ctxTimeout, parsedId); err != nil {
		statusCode := http.StatusInternalServerError

		if err == ErrorUserNotFound {
			statusCode = http.StatusNotFound
		}

		ctx.JSON(statusCode, gin.H{
			"sucess": false,
			"data":   nil,
			"error":  err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{
		"sucess": true,
		"data":   nil,
		"error":  nil,
	})
}
