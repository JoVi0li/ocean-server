package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/JoVi0li/ocean-server/internal/shared/database"
	"github.com/JoVi0li/ocean-server/internal/shared"
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

func SignUp(ctx *gin.Context) {
	var user User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"sucess": false,
			"data":   nil,
			"error":  err.Error(),
		})
		return
	}

	_, err := service.FindByEmail(ctx, user.Email)
	if err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"sucess": false,
			"data":   nil,
			"error":  shared.ErrorCredentialsInvalid.Error(),
		})
		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	sucess, err := service.Create(ctxTimeout, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"sucess": false,
			"data":   nil,
			"error":  err.Error(),
		})
		return
	}

	userInfoToken := &shared.UserInfoToken{
		ID:       sucess.ID.String(),
		Username: sucess.Username,
		Email:    sucess.Email,
	}

	token, err := shared.NewToken(*userInfoToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"sucess": false,
			"data":   nil,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"sucess": true,
		"data":   token,
		"error":  nil,
	})
}

func SignIn(ctx *gin.Context) {
	var infos SignInDTO

	if err := ctx.BindJSON(&infos); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"sucess": false,
			"data":   nil,
			"error":  err.Error(),
		})
		return
	}

	user, err := service.FindByEmail(ctx, infos.Email)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"sucess": false,
			"data":   nil,
			"error":  shared.ErrorCredentialsInvalid.Error(),
		})
		return
	}

	if err := shared.CheckPassword(user.Password, infos.Password); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"sucess": false,
			"data":   nil,
			"error":  shared.ErrorCredentialsInvalid.Error(),
		})
		return
	}

	userInfoToken := &shared.UserInfoToken{
		ID:       user.ID.String(),
		Username: user.Username,
		Email:    user.Email,
	}

	token, err := shared.NewToken(*userInfoToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"sucess": false,
			"data":   nil,
			"error":  shared.ErrorCredentialsInvalid.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"sucess": true,
		"data":   token,
		"error":  nil,
	})
}

func DeleteAccount(ctx *gin.Context) {
	token, err := shared.GetToken(ctx)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"sucess": false,
			"data":   nil,
			"error":  err.Error(),
		})
		return
	}

	decodedTk, err := shared.DecodeTokenClaims(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"sucess": false,
			"data":   nil,
			"error":  err.Error(),
		})
		return
	}

	parsedId, err := uuid.Parse(decodedTk.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"sucess": false,
			"data":   nil,
			"error":  shared.ErrorIdInvalid.Error(),
		})
		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := service.DeleteById(ctxTimeout, parsedId); err != nil {
		statusCode := http.StatusInternalServerError
		if err == shared.ErrorUserNotFound {
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
