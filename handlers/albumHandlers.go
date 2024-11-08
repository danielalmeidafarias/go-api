package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/danielalmeidafarias/go-api/models"

	albums "github.com/danielalmeidafarias/go-api/db"

	"github.com/danielalmeidafarias/go-api/helpers"
)

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums.Albums)
}

func AddAlbum(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	newAlbum.ID = helpers.GetLastId()

	if len(newAlbum.Artist) == 0 {
		data := gin.H{"message": `required field "artist" is missing`}
		c.IndentedJSON(http.StatusBadRequest, data)
		return
	}

	if len(newAlbum.Title) == 0 {
		data := gin.H{"message": `required field "title" is missing`}
		c.IndentedJSON(http.StatusBadRequest, data)
		return
	}

	if newAlbum.Price < 0 {
		data := gin.H{"message": "price must be positive"}
		c.IndentedJSON(http.StatusBadRequest, data)
		return
	}

	albums.Albums = append(albums.Albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, albums.Albums)
}

func GetOneAlbum(c *gin.Context) {
	id := c.Param("id")

	for i := range albums.Albums {
		if albums.Albums[i].ID == id {
			c.IndentedJSON(http.StatusOK, albums.Albums[i])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": fmt.Sprintf("No user with %v ID was found", id),
	})
}
