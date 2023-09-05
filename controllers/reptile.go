package controllers

import (
	"ReptileApi/models"
	"fmt"
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
	// var delete_item models.Reptile
	fmt.Println("este es el id:", c.Param("data"))
}

func PostReptiles(c *gin.Context) {
	var reptile models.Reptile

	if err := c.BindJSON(&reptile); err != nil {
		log.Fatal(err)
		return
	}
	ReptilList = append(ReptilList, reptile)
	c.IndentedJSON(http.StatusCreated, reptile)
}
