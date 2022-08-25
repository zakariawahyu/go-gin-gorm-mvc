package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zakariawahyu/go-gin-gorm-mvc/entity"
	"github.com/zakariawahyu/go-gin-gorm-mvc/models"
	"net/http"
	"time"
)

func GetCustomers(c *gin.Context) {
	var customers []entity.Customer
	err := models.GetAllCustomers(&customers)

	if err != nil {
		c.JSONP(http.StatusNotFound, gin.H{
			"message": "Customer nil",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"customers": customers,
		})
	}
}

func CreateCustomers(c *gin.Context) {
	var customer entity.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "cannot handle request",
		})
		return
	}

	customer.CreatedAt = time.Now()
	customer.UpdateAt = time.Now()
	err := models.CreateCustomers(&customer)
	
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"message":  "data created successfully",
			"customer": customer,
		})
		return
	}
}
