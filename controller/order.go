package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zakariawahyu/go-gin-gorm-mvc/entity"
	"github.com/zakariawahyu/go-gin-gorm-mvc/models"
	"net/http"
)

func GetOrder(c *gin.Context) {
	var order []entity.Order
	err := models.GetAllOrder(&order)

	if err != nil {
		c.JSONP(http.StatusNotFound, gin.H{
			"message": "Order nil",
		})
	} else {
		c.JSON(http.StatusOK, order)
	}
}
