package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	listenAddr := ":9999"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		// FUNCTIONS_CUSTOMHANDLER_PORT is set by azure func runtime
		listenAddr = ":" + val
	}
	r.GET("/api/albums", GetAlbums)
	r.GET("/api/albums/:id", GetAlbumByID)
	r.POST("/api/albums", PostAlbums)
	r.Run(listenAddr) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// Album represents data about a record album.
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "4", Title: "fdsfdsf", Artist: "aaa", Price: 39.99},
	{ID: "5", Title: "SfsdfdsfsdfsBrown", Artist: "aaaaaaaa", Price: 39.99},
	{ID: "6", Title: "Sardfsfdsfsdfdsfsdfdsfsdfdsfsdown", Artist: "aaaaaaaaaaaaaa", Price: 99.99},
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// PostAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
	var newAlbum Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// GetAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
