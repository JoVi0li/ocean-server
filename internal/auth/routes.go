package auth

import "github.com/gin-gonic/gin"

func SetRoutes(g *gin.Engine) {
	g.POST("/auth/sigin", SignIn)
	g.POST("/auth/signup", SignUp)
}