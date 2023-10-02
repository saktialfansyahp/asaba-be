package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

	// r.POST("/api/login", authcontroller.Login)
	// r.POST("/api/register", authcontroller.Register)
	// r.POST("/api/role", authcontroller.Role)
	// r.GET("/api/role", authcontroller.GetRole)

	// r.GET("/api/kategori", kategoricontroller.Index)
	// r.POST("/api/kategori", kategoricontroller.Create)

	// r.GET("/api/status", statuscontroller.Index)
	// r.POST("/api/status", statuscontroller.Create)

	// r.GET("/api/sale", produkcontroller.Sale)

	// r.GET("/api/produk", produkcontroller.Index)
	// r.POST("/api/produk", produkcontroller.Create)
	// r.GET("/api/produk/:id", func(ctx *gin.Context) {
	// 	produkcontroller.ById(ctx, ctx.Param("id"))
	// })
	// r.PUT("/api/produk/:id", func(ctx *gin.Context) {
	// 	produkcontroller.Edit(ctx, ctx.Param("id"))
	// })
	// r.DELETE("/api/produk/:id", func(ctx *gin.Context) {
	// 	produkcontroller.Delete(ctx, ctx.Param("id"))
	// })

	r.Run()
}