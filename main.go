package main

import (
	"ReptileApi/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	controllers.ConnectDatabase()

	router := gin.Default()
	router.GET("/reptiles", controllers.GetReptiles)
	router.POST("/reptiles", controllers.PostReptiles)
	router.DELETE("/reptiles", controllers.DeleteReptiles)

	router.Run("localhost:8080")
}
