package main

import (
	"context"
	"log"
	"os"

	"github.com/JoVi0li/ocean-server/internal/auth"
	"github.com/JoVi0li/ocean-server/internal/database"
	"github.com/JoVi0li/ocean-server/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	/// Setup server
	server := gin.Default()

	/// Setup environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	/// Setup database
	_, err := database.NewConnection(context.Background(), os.Getenv("POSTGRES_CONN_STR"))
	if err != nil {
		log.Fatal(err)
	}

	/// Setup user
	user.Configure()
	user.SetRoutes(server, auth.AuthMidd())

	/// Setup auth
	auth.Configure()
	auth.SetRoutes(server)

	/// Run server
	server.Run(":3000")
}
