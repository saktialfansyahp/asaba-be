package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saktialfansyahp/asaba-be.git/models"
)

func Index(c *gin.Context) {
	var barang []models.Barang

	if err := models.DB.Find(&barang).Error; err != nil {
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

    for _, input := range barang {
        if err := models.DB.Create(&input).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
            return
        }
    }

	c.JSON(http.StatusCreated, gin.H{"error": "0", "data": &barang})
}

func Stok(c *gin.Context) {
	var stoks []models.Stok

	if err := c.ShouldBindJSON(&stoks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	var Success = true

	var newStok []models.Barang

	for _, stok := range stoks {
		var existingBarang models.Barang

		if err := models.DB.Where("code = ?", stok.Code).First(&existingBarang).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Not Found", "message": "Barang not found"})
			Success = false 
			break 
		}

		if stok.Tipe != "masuk" && stok.Tipe != "keluar" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Invalid request type"})
			Success = false
			break
		}

		newStok = append(newStok, existingBarang)
	}

	if Success {
		var success = true
		for i, stok := range stoks {
			if stok.Tipe == "masuk" {
				newStok[i].Jumlah += stok.Jumlah
			} else if stok.Tipe == "keluar" {
				if newStok[i].Jumlah < stok.Jumlah {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Not enough stock"})
					success = false
					break
				}
				newStok[i].Jumlah -= stok.Jumlah
			}
		}

		if success {
			for i := range newStok {
				if err := models.DB.Save(&newStok[i]).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
					return
				}
			}
			for i, stok := range stoks {
				history := models.Stok{
					Code: newStok[i].Code,
					Tipe: stok.Tipe,
					Jumlah: stok.Jumlah,
				}

				if err := models.DB.Create(&history).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
					return
				}
			}

			c.JSON(http.StatusOK, gin.H{"message": "Stok barang berhasil diupdate"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Data tidak valid, penyimpanan dibatalkan"})
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Data tidak valid, penyimpanan dibatalkan"})
	}
}

func StokIndex(c *gin.Context) {
	var barang []models.Stok

	if err := models.DB.Find(&barang).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": 1})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": 0, "data": barang})
}