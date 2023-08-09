package main

import (
	"simpleapi_go/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const dbuser = "root"
const dbpass = ""
const dbname = "gin"

func main() {
	router := gin.Default()

	productService := models.NewProductService(dbuser, dbpass, dbname)

	router.GET("/products", productService.GetProducts)
	router.GET("/product/:code", productService.GetProduct)
	router.POST("/products", productService.AddProduct)

	router.Run("localhost:8083")
}
