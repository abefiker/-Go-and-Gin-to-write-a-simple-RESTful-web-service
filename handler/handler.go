package handler

import (
	"net/http"

	"example/web-service-gin/data"
	"example/web-service-gin/model"

	"github.com/gin-gonic/gin"
)

// GetAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, data.Albums)
}

// PostAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
    var newAlbum model.Album

    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    data.Albums = append(data.Albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

// GetAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {
    id := c.Param("id")

    for _, a := range data.Albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
