package controllers

import (
	"ReptileApi/db"
	"ReptileApi/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// var ReptilList = []models.Reptile{
// 	{Name: "Mohave", RegularPrice: 2500, Price: 1999, AgeCategory: models.Cria, Description: "Bonita mohave pastel", Genre: models.HEMBRA},
// 	{Name: "Leopard Het-Clown", RegularPrice: 2000, Price: 1000, AgeCategory: models.Cria, Description: "Leopard", Genre: models.MACHO},
// 	{Name: "Spotnose probable Blade 100% Het-Clown", RegularPrice: 9500, Price: 9200, AgeCategory: models.Cria, Description: "Bonita Leopard Het-Clown", Genre: models.HEMBRA},
// 	{Name: "Super Enchi Orange Dream Pastel ", RegularPrice: 2500, Price: 1999, AgeCategory: models.Cria, Description: "Bonita Het Clown", Genre: models.MACHO},
// }

func GetReptiles(c *gin.Context) {
	var newlist []models.Reptile
	db.RetriveReptiles(&newlist)
	log.Println(newlist)
	// c.IndentedJSON(http.StatusOK, ReptilList)
}

func DeleteReptiles(c *gin.Context) {

	fmt.Println("este es el id:", c.Param("data"))
}

func PostReptiles(c *gin.Context) {
	var reptile models.Reptile

	if err := c.BindJSON(&reptile); err != nil {
		log.Fatal(err)
		return
	}
	// ReptilList = append(ReptilList, reptile)
	c.IndentedJSON(http.StatusCreated, reptile)
}
