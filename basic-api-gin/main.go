package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(cx *gin.Context) {
	cx.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(cx *gin.Context) {
	id := cx.Param("id")

	for _, album := range albums {
		if album.ID == id {
			cx.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	cx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbums(cx *gin.Context) {
	var newAlbum album
	if err := cx.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	cx.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
