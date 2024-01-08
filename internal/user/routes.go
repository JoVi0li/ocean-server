package user

import (
	"github.com/gin-gonic/gin"
)

func SetRoutes (g *gin.Engine, authMidd gin.HandlerFunc) {
	group := g.Group("/user")
	group.Use(authMidd)
	group.GET("/", GetUsers)
	group.DELETE("/", DeleteUsers)
}