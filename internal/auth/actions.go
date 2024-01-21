package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/JoVi0li/ocean-server/internal/database"
	"github.com/JoVi0li/ocean-server/internal/user"
	"github.com/JoVi0li/ocean-server/internal/util"
	"github.com/gin-gonic/gin"
)

var service user.Service

func Configure() {
	service = user.Service{
		Repository: &user.RepositoryPostgres{
			Connection: database.Conn,
		},
	}
}

func SignUp(ctx *gin.Context) {
	var user user.User

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
			"error":  ErrorCredentialsInvalid.Error(),
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

	userInfoToken := &util.UserInfoToken{
		ID:       sucess.ID.String(),
		Username: sucess.Username,
		Email:    sucess.Email,
	}

	token, err := util.NewToken(*userInfoToken)

	if err != nil {
		ctx.JSON(http.StatusCreated, gin.H{
			"sucess": true,
			"data":   sucess,
			"error":  nil,
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"sucess": true,
		"data":   token,
		"error":  nil,
	})
}

func SignIn(ctx *gin.Context) {
	var infos SignInInfos

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
			"error":  ErrorCredentialsInvalid.Error(),
		})

		return
	}

	if err := util.CheckPassword(user.Password, infos.Password); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"sucess": false,
			"data":   nil,
			"error":  ErrorCredentialsInvalid.Error(),
		})

		return
	}

	userInfoToken := &util.UserInfoToken{
		ID:       user.ID.String(),
		Username: user.Username,
		Email:    user.Email,
	}

	token, err := util.NewToken(*userInfoToken)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"sucess": false,
			"data":   nil,
			"error":  ErrorCredentialsInvalid.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"sucess": true,
		"data":   token,
		"error":  nil,
	})
}
