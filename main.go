package main

import (
	"ReptileApi/controllers"
	"ReptileApi/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectDb()

	router := gin.Default()
	router.GET("/reptiles", controllers.GetReptiles)
	router.POST("/reptiles", controllers.PostReptiles)
	router.DELETE("/reptiles", controllers.DeleteReptiles)

	router.Run("localhost:8080")
}
