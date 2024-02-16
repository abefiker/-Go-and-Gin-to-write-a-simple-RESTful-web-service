package router

import (
    "github.com/gin-gonic/gin"
    "example/web-service-gin/handler"
)

// SetupRouter sets up the router and defines routes.
func SetupRouter() *gin.Engine {
    router := gin.Default()

    router.GET("/albums", handler.GetAlbums)
    router.GET("/albums/:id", handler.GetAlbumByID)
    router.POST("/albums", handler.PostAlbums)

    return router
}
