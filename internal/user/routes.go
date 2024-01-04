package user

import "github.com/gin-gonic/gin"

func SetRoutes (g *gin.Engine) {
	g.GET("/users/:id", GetUsers)
	g.DELETE("/users/:id", DeleteUsers)
}