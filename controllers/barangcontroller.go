package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saktialfansyahp/asaba-be.git/models"
)

func Index(c *gin.Context) {
	var barang []models.Barang

	if err := models.DB.Preload("Kategori").Preload("Status").Find(&barang).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": 1})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": 0, "data": barang})
}

func ById(c *gin.Context, id string) {
	var barang models.Barang

	if err := models.DB.Preload("Kategori").Preload("Status").Where("id = ?", id).First(&barang).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": 1})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": 0, "data": barang})
}

func Create(c *gin.Context) {
	var barang []models.Barang

	if err := c.ShouldBindJSON(&barang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	for _, input := range barang {
        var existingBarang models.Barang
        if err := models.DB.Where("code = ?", input.Code).First(&existingBarang).Error; err == nil {
            c.JSON(http.StatusConflict, gin.H{"error": "Conflict", "message": "Barang with the same code already exists"})
            return
        }
    }

    // Simpan setiap barang ke database
    for _, input := range barang {
        if err := models.DB.Create(&input).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
            return
        }
    }

	if err := models.DB.Create(&barang).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "1"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"error": "0", "data": &barang})
}

func Edit(c *gin.Context, id string) {
	var barang models.Barang
	if err := c.ShouldBindJSON(&barang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "1"})
		return
	}

	// Cek apakah barang dengan product_id tersebut ada dalam database
	var existingProduct models.Barang
	if err := models.DB.Where("id = ?", id).First(&existingProduct).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "1"})
		return
	}

	if err := models.DB.Save(&existingProduct).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "1"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": "0", "data": &existingProduct})
}

func Delete(c *gin.Context, id string) {
	var existingProduct models.Barang
	if err := models.DB.Where("id = ?", id).First(&existingProduct).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "1"})
		return
	}

	if err := models.DB.Delete(&existingProduct).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "1"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"error": "0", "data": &existingProduct})
}