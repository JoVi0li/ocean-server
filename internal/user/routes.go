package user

import "github.com/gin-gonic/gin"

func SetRoutes (g *gin.Engine) {
	g.POST("/users", CreateUsers)
	g.GET("/users/:id", GetUsers)
	g.DELETE("/users/:id", DeleteUsers)
}