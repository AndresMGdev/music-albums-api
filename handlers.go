package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


var albums = []album{
	{Id: 1, Title: "Sobre El Amor Y Sus Efectos Secundarios", Artist: "Morat", Year: 2017},
	{Id: 2, Title: "Balas Perdidas", Artist: "Morat", Year: 2018},
	{Id: 3, Title: "¿A donde Vamos?", Artist: "Morat", Year: 2021},
	{Id: 4, Title: "SI AYER FUERA HOY", Artist: "Morat", Year: 2022},
	{Id: 5, Title: "1989 (Taylor's Version)", Artist: "Taylor Swift", Year: 2023},
	{Id: 6, Title: "Fearless", Artist: "Taylor Swift", Year: 2008},
	{Id: 7, Title: "Lover", Artist: "Taylor Swift", Year: 2019},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, albums)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")
	iid, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	for _, album := range albums {
		if album.Id == iid {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

func deleteAlbum(c *gin.Context) {
	id := c.Param("id")
	iid, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	for i, album := range albums {
		if album.Id == iid {
			albums = append(albums[:i], albums[i+1:]...)

			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Album not found"})
}
