package main

import (
	"ReptileApi/controllers"
	"ReptileApi/db"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Connectdb()
	db.Connectdb()
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://localhost*"},
		AllowMethods:     []string{"GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://localhost:8080/*"
		},
		MaxAge: 12 * time.Hour,
	}))
	router.GET("/reptiles", controllers.GetReptiles)
	router.POST("/reptiles", controllers.PostReptiles)
	router.DELETE("/reptiles", controllers.DeleteReptiles)

	router.Run("localhost:8080")
}
