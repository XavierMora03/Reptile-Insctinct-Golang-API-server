package models

// const MAX_PHOTOS_PER_REPTILE_PRODUCT int = 10
import (
	"github.com/Rhymond/go-money"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func trashFUnciont() {
	a := money.New(100, money.MXN)
	log.Fatal(a)
}

type Reptile struct {
	ID           uint64 `json:"id"`
	Name         string `json:"name"`
	RegularPrice uint64 `json:"regularPrice"`
	Price        uint64 `json:"price"`
	AgeCategory  string `json:"age"`
	pictures     *ReptilePictures
	Description  string `json:"description"`
	Genre        string `json:"genre"`
}

type ReptilePictures struct {
}

var ReptilList = []Reptile{
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
	var newAlbum Reptile

	if err := c.BindJSON(&newAlbum); err != nil {
		log.Fatal(err)
		return
	}
	ReptilList = append(ReptilList, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
