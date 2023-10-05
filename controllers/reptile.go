package controllers

import (
	_ "ReptileApi/db"
	"ReptileApi/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var ReptilList = []models.Reptile{
	{Name: "Mohave", RegularPrice: "2500", Price: "1999", AgeCategory: models.Cria, Description: "bonita mohave pastel", Genre: models.HEMBRA},
	{Name: "leopard het-clown", RegularPrice: "2000", Price: "1000", AgeCategory: models.Cria, Description: "leopard", Genre: models.MACHO},
	{Name: "spotnose probable blade 100% het-clown", RegularPrice: "9500", Price: "9200", AgeCategory: models.Cria, Description: "bonita leopard het-clown", Genre: models.HEMBRA},
	{Name: "super enchi orange dream pastel ", RegularPrice: "2500", Price: "1999", AgeCategory: models.Cria, Description: "bonita het clown", Genre: models.MACHO},
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
func GetReptiles(c *gin.Context) {
	//ReptilList = nil
	//db.RetriveReptiles(&ReptilList)
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

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
