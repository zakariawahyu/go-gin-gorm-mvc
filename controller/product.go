package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zakariawahyu/go-gin-gorm-mvc/entity"
	"github.com/zakariawahyu/go-gin-gorm-mvc/models"
	"net/http"
	"strconv"
	"time"
)

func GetProduct(c *gin.Context) {
	var product []entity.Product
	err := models.GetAllProduct(&product)

	if err != nil {
		c.JSONP(http.StatusNotFound, gin.H{
			"message": "Product nil",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"products": product,
		})
	}
}

func CreateProduct(c *gin.Context) {
	var product entity.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "cannot handle request",
		})
		return
	}

	product.CreatedAt = time.Now()
	product.UpdateAt = time.Now()
	err := models.CreateProduct(&product)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"message": "data created successfully",
			"product": product,
		})
		return
	}
}

func ShowProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	var product entity.Product

	err := models.GetProductByID(&product, id)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"product": product,
		})
		return
	}

}
