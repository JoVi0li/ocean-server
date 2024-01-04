package auth

import "github.com/gin-gonic/gin"

func SetRoutes(g *gin.Engine) {
	g.POST("/sigin", SignIn)
	g.POST("/signup", SignUp)
}