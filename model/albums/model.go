package albums

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Album represents data about a record album.
type Album struct {
	ID      string  `json:"id"`
	Title   string  `json:"title"`
	Artist  string  `json:"artist"`
	Price   float64 `json:"price"`
	Summary float64 `json:"summary"`
}

// albums slice to seed record album data.
var albums = []Album{
	{ID: "1", Title: "RED Train", Artist: "John Coltrane", Price: 56.99, Summary: 1000},
	{ID: "2", Title: "GREEN Train", Artist: "John Coltrane", Price: 56.99, Summary: 1000},
	{ID: "3", Title: "GREEN Train", Artist: "John Coltrane", Price: 56.99, Summary: 1000},
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
