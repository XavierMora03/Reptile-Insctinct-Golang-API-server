package controllers

import (
	"ReptileApi/db"
	"ReptileApi/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var ReptilList = []models.Reptile{
	{Name: "Mohave", RegularPrice: "2500", Price: "1999", AgeCategory: models.Cria, Description: "bonita mohave pastel", Genre: models.HEMBRA},
	{Name: "leopard het-clown", RegularPrice: "2000", Price: "1000", AgeCategory: models.Cria, Description: "leopard", Genre: models.MACHO},
	{Name: "spotnose probable blade 100% het-clown", RegularPrice: "9500", Price: "9200", AgeCategory: models.Cria, Description: "bonita leopard het-clown", Genre: models.HEMBRA},
	{Name: "super enchi orange dream pastel ", RegularPrice: "2500", Price: "1999", AgeCategory: models.Cria, Description: "bonita het clown", Genre: models.MACHO},
}

func GetReptiles(c *gin.Context) {
	ReptilList = nil
	db.RetriveReptiles(&ReptilList)
	c.IndentedJSON(http.StatusOK, ReptilList)
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
