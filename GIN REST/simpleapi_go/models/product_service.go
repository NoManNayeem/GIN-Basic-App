package models

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductService struct {
	productDAO *ProductDAO
}

func NewProductService(dbuser, dbpass, dbname string) *ProductService {
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)
	if err != nil {
		panic(err.Error())
	}
	return &ProductService{
		productDAO: NewProductDAO(db),
	}
}

func (service *ProductService) GetProducts(c *gin.Context) {
	products := service.productDAO.GetProducts()
	if products == nil || len(products) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, products)
	}
}

func (service *ProductService) GetProduct(c *gin.Context) {
	code := c.Param("code")
	product := service.productDAO.GetProduct(code)
	if product == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, product)
	}
}

func (service *ProductService) AddProduct(c *gin.Context) {
	var prod Product
	if err := c.BindJSON(&prod); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		service.productDAO.AddProduct(prod)
		c.IndentedJSON(http.StatusCreated, prod)
	}
}
