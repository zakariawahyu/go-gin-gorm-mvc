package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zakariawahyu/go-gin-gorm-mvc/entity"
	"github.com/zakariawahyu/go-gin-gorm-mvc/models"
	"net/http"
)

func GetProduct(c *gin.Context) {
	var product []entity.Product
	err := models.GetAllProduct(&product)

	if err != nil {
		c.JSONP(http.StatusNotFound, gin.H{
			"message": "Customer nil",
		})
	}
	c.JSON(http.StatusOK, product)
}
