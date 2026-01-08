package handlers

import (
	"net/http"

	"bioskop-app-sb-go-batch-73-Agil/config"
	"bioskop-app-sb-go-batch-73-Agil/models"
	"github.com/gin-gonic/gin"
)

func CreateBioskop(c *gin.Context) {
	var bioskops []models.Bioskop

	if err := c.ShouldBindJSON(&bioskops); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Format JSON harus berupa array",
		})
		return
	}

	for _, b := range bioskops {
		if b.Nama == "" || b.Lokasi == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Nama dan Lokasi tidak boleh kosong",
			})
			return
		}
	}

	config.DB.Create(&bioskops)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Data bioskop berhasil ditambahkan",
		"data":    bioskops,
	})
}


func GetAllBioskop(c *gin.Context) {
	var bioskops []models.Bioskop

	result := config.DB.Find(&bioskops)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}

	c.JSON(http.StatusOK, bioskops)
}

func GetBioskopByID(c *gin.Context) {
	var bioskop models.Bioskop
	id := c.Param("id")

	result := config.DB.First(&bioskop, id)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bioskop tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, bioskop)
}

func UpdateBioskop(c *gin.Context) {
	var bioskop models.Bioskop
	id := c.Param("id")

	if err := config.DB.First(&bioskop, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bioskop tidak ditemukan"})
		return
	}

	var input models.Bioskop
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON tidak valid"})
		return
	}

	if input.Nama == "" || input.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
		return
	}

	bioskop.Nama = input.Nama
	bioskop.Lokasi = input.Lokasi
	bioskop.Rating = input.Rating

	config.DB.Save(&bioskop)

	c.JSON(http.StatusOK, gin.H{
		"message": "Bioskop berhasil diperbarui",
		"data":    bioskop,
	})
}

func DeleteBioskop(c *gin.Context) {
	var bioskop models.Bioskop
	id := c.Param("id")

	if err := config.DB.First(&bioskop, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bioskop tidak ditemukan"})
		return
	}

	config.DB.Delete(&bioskop)

	c.JSON(http.StatusOK, gin.H{
		"message": "Bioskop berhasil dihapus",
	})
}

