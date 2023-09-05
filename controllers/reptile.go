package controllers

import (
	"ReptileApi/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var ReptilList = []models.Reptile{
	{Name: "Mohave", RegularPrice: 2500, Price: 1999, AgeCategory: "Cria", Description: "Bonita mohave pastel", Genre: "Hembra"},
	{Name: "Leopard Het-Clown", RegularPrice: 2000, Price: 1000, AgeCategory: "Cria", Description: "Leopard", Genre: "Macho"},
	{Name: "Spotnose probable Blade 100% Het-Clown", RegularPrice: 9500, Price: 9200, AgeCategory: "Cria", Description: "Bonita Leopard Het-Clown", Genre: "Hembra"},
	{Name: "Super Enchi Orange Dream Pastel ", RegularPrice: 2500, Price: 1999, AgeCategory: "Cria", Description: "Bonita Het Clown", Genre: "Hembra"},
}

func GetReptiles(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, ReptilList)
}

func DeleteReptiles(c *gin.Context) {
	// var id_delete uint64 = c.GetHeader()
}

func PostReptiles(c *gin.Context) {
	var newAlbum models.Reptile

	if err := c.BindJSON(&newAlbum); err != nil {
		log.Fatal(err)
		return
	}
	ReptilList = append(ReptilList, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
