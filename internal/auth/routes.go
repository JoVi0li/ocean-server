package auth

import "github.com/gin-gonic/gin"

func SetRoutes(g *gin.Engine) {
	group := g.Group("/auth")
	group.GET("/signin", SignIn)
	group.DELETE("/signup", SignUp)
}