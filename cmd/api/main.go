package main

import (
	"context"
	"log"
	"os"

	"github.com/JoVi0li/ocean-server/internal/modules/auth"
	"github.com/JoVi0li/ocean-server/internal/shared/database"
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

	/// Setup auth
	auth.Configure()
	auth.SetRoutes(server)

	/// Run server
	server.Run(":3000")
}
