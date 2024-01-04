package main

import (
	"github.com/JoVi0li/ocean-server/internal/user"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	user.Configure()
	user.SetRoutes(server)
	server.Run(":3000")
}