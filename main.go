package main

import (
	"ReptileApi/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Print("SERVER RUNNING")
	router := gin.Default()
	router.GET("/reptiles", controllers.GetReptiles)
	router.POST("/reptiles", controllers.PostReptiles)

	router.Run("localhost:8080")
}
