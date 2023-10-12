package main

import (
	"ReptileApi/controllers"
	"ReptileApi/db"
	"ReptileApi/imageManagment"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	imagemanagment.SaveImage(1, []byte{5, 56, 5, 4, 5, 6, 5})

	return
	db.Connectdb()
	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/reptiles", controllers.GetReptiles)
	router.POST("/reptiles", controllers.PostReptiles)
	router.DELETE("/reptiles", controllers.DeleteReptiles)

	router.Run("localhost:8080")
}
