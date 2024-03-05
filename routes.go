package main

import "github.com/gin-gonic/gin"

func setupRoutes(router *gin.Engine) {
	router.GET("api/albums", getAlbums)
	router.POST("api/albums", postAlbums)
	router.GET("api/albums/album/:id", getAlbumById)
	router.GET("api/albums/album/delete/:id", deleteAlbum)
}