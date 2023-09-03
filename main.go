package main

import (
	// "encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
type mensaje struct {
	Nombre string `json:"nombre"`
	Amor   string `json:"amor_de_mi_vida"`
	Msj    string `json:"teamomucho"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// get albumns function
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	fmt.Print("SERVER RUNNING")
	router := gin.Default()
	router.GET("/albumns", getAlbums)

	router.Run("localhost:8080")
}