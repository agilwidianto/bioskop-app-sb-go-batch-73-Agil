package main

import (
	"os"

	"bioskop-app-sb-go-batch-73-Agil/config"
	"bioskop-app-sb-go-batch-73-Agil/handlers"
	"bioskop-app-sb-go-batch-73-Agil/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Bioskop{})

	r.POST("/bioskop", handlers.CreateBioskop)
	r.GET("/bioskop", handlers.GetAllBioskop)
	r.GET("/bioskop/:id", handlers.GetBioskopByID)
	r.PUT("/bioskop/:id", handlers.UpdateBioskop)
	r.DELETE("/bioskop/:id", handlers.DeleteBioskop)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}