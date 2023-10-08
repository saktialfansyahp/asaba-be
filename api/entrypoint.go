package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/saktialfansyahp/asaba-be.git/controllers"
	"github.com/saktialfansyahp/asaba-be.git/models"
)

var (
	app *gin.Engine
)

func router(r *gin.RouterGroup){
	models.ConnectDatabase()

	r.GET("api/barang", controllers.Index)
	r.POST("api/barang", controllers.Create)
	r.POST("api/stok", controllers.Stok)
}

func init() {
	app = gin.New()

	app.NoRoute(func(c *gin.Context) {
		sb := &strings.Builder{}
		sb.WriteString("routing err: no route, try this:\n")
		for _, v := range app.Routes() {
			sb.WriteString(fmt.Sprintf("%s %s\n", v.Method, v.Path))
		}
		c.String(http.StatusBadRequest, sb.String())
	})

	r := app.Group("/")

	router(r)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}