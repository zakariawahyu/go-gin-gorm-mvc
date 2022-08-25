package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zakariawahyu/go-gin-gorm-mvc/entity"
	"github.com/zakariawahyu/go-gin-gorm-mvc/models"
	"net/http"
)

func GetOrderDetail(c *gin.Context) {
	var orderDetail []entity.OrderDetail
	err := models.GetAllOrderDetail(&orderDetail)

	if err != nil {
		c.JSONP(http.StatusNotFound, gin.H{
			"message": "Order detail nil",
		})
	} else {
		c.JSON(http.StatusOK, orderDetail)
	}
}
