package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zakariawahyu/go-gin-gorm-mvc/entity"
	"github.com/zakariawahyu/go-gin-gorm-mvc/models"
	"net/http"
)

func GetCustomers(c *gin.Context) {
	var customers []entity.Customer
	err := models.GetAllCustomers(&customers)

	if err != nil {
		c.JSONP(http.StatusNotFound, gin.H{
			"message": "Customer nil",
		})
	} else {
		c.JSON(http.StatusOK, customers)
	}
}
