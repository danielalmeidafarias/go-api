package main

import (
	"os"

	"github.com/danielalmeidafarias/go-api/handlers"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	router := gin.Default()

	router.GET("/albums", handlers.GetAlbums)

	router.GET("/albums/:id", handlers.GetOneAlbum)

	router.POST("/albums", handlers.AddAlbum)

	router.Run(os.Getenv("API_URL"))
}
