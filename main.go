package main

import (
	"ReptileApi/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	controllers.ConnectDatabase()

	fmt.Print("SERVER RUNNING")
	router := gin.Default()
	router.GET("/reptiles", controllers.GetReptiles)
	router.POST("/reptiles", controllers.PostReptiles)
	router.DELETE("/reptiles", controllers.DeleteReptiles)

	router.Run("localhost:8080")
}
