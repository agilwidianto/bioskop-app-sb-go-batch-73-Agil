package handlers

import (
	"net/http"

	"bioskop-app-sb-go-batch-73-Agil/config"
	"bioskop-app-sb-go-batch-73-Agil/models"
	"github.com/gin-gonic/gin"
)

func CreateBioskop(c *gin.Context) {
	var bioskop models.Bioskop

	if err := c.ShouldBindJSON(&bioskop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Format JSON tidak valid",
		})
		return
	}

	// Validasi
	if bioskop.Nama == "" || bioskop.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Nama dan Lokasi tidak boleh kosong",
		})
		return
	}

	config.DB.Create(&bioskop)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Bioskop berhasil ditambahkan",
		"data":    bioskop,
	})
}
