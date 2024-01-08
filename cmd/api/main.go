package main

import (
	"github.com/JoVi0li/ocean-server/internal/auth"
	"github.com/JoVi0li/ocean-server/internal/user"
	"github.com/gin-gonic/gin"
)

func main() {
	/// Setup server
	server := gin.Default()

	/// Setup user
	user.Configure()
	user.SetRoutes(server, auth.AuthMidd())

	/// Setup auth
	auth.Configure()
	auth.SetRoutes(server)

	/// Run server
	server.Run(":3000")
}