package auth

import "github.com/gin-gonic/gin"

func SetRoutes(g *gin.Engine) {
	group := g.Group("/auth")
	group.POST("/signin", SignIn)
	group.POST("/signup", SignUp)
}
