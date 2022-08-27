package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zakariawahyu/go-gin-gorm-mvc/entity"
	"github.com/zakariawahyu/go-gin-gorm-mvc/models"
	"net/http"
	"time"
)

func GetOrder(c *gin.Context) {
	var order []entity.Order
	err := models.GetAllOrder(&order)

	if err != nil {
		c.JSONP(http.StatusNotFound, gin.H{
			"message": "Order nil",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"orders": order,
		})
	}
}

func CreateOrder(c *gin.Context) {
	var order entity.Order
	var orderDetail entity.OrderDetail
	var product entity.Product

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "cannot handle request",
			"err":     order,
		})
		return
	}

	order.CreatedAt = time.Now()
	order.UpdateAt = time.Now()
	order.Status = "created"
	err := models.CreateOrder(&order)

	for _, detailOrder := range order.OrderDetailCustomer {
		_ = models.ShowProduct(&product, detailOrder.ID)
		orderDetail.OrderID = order.ID
		orderDetail.ProductID = detailOrder.ID
		orderDetail.Qty = detailOrder.Qty
		orderDetail.Amount = detailOrder.Qty * product.Price
		orderDetail.CreatedAt = time.Now()
		orderDetail.UpdateAt = time.Now()
		models.CreateOrderDetail(&orderDetail)
	}

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"message": "data created successfully",
			"order":   order,
		})
	}
}
