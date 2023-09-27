package main

import (
	"ReptileApi/controllers"
	"ReptileApi/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Connectdb()

	router := gin.Default()
	router.Use(cors.New(cors.DefaultConfig()))
	router.GET("/reptiles", controllers.GetReptiles)
	router.POST("/reptiles", controllers.PostReptiles)
	router.DELETE("/reptiles", controllers.DeleteReptiles)

	router.Run("localhost:8080")
}
