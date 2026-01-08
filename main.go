package main

import (
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

	r.Run(":8080")
}
