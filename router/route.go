package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saktialfansyahp/asaba-be.git/controllers"
	"github.com/saktialfansyahp/asaba-be.git/models"
)

func DefineRoutes() {
	r := gin.Default()
	models.ConnectDatabase()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
			c.JSON(http.StatusOK, gin.H{"message": "Preflight request successful"})
			c.Abort()
			return
		}

		if c.Request.Method == "POST" {
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")
		} else if c.Request.Method == "GET" {
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
		} else if c.Request.Method == "PUT" {
			c.Writer.Header().Set("Access-Control-Allow-Methods", "PUT")
		} else if c.Request.Method == "DELETE" {
			c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE")
		}

		c.Next()
	})

	r.GET("api/barang", controllers.Index)
	r.POST("api/barang", controllers.Create)
	r.POST("api/stok", controllers.Stok)

	r.Run()
}