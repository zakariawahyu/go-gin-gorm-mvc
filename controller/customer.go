package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zakariawahyu/go-gin-gorm-mvc/config"
	"github.com/zakariawahyu/go-gin-gorm-mvc/entity"
	"github.com/zakariawahyu/go-gin-gorm-mvc/models"
	"net/http"
	"strconv"
	"time"
)

func GetCustomers(c *gin.Context) {
	var customers []entity.CustomerWithoutOrder
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

func GetCustomersWithOrder(c *gin.Context) {
	var customers []entity.Customer
	err := models.GetAllCustomersWithOrder(&customers)

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
	var customer entity.CustomerResponse
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "cannot handle request",
		})
		return
	}

	customer.IsActive = true
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

func ShowCustomer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	var customer entity.CustomerWithoutOrder

	err := models.ShowCustomer(&customer, id)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	} else {
		if customer.Username == "" {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Customer not found",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"customer": customer,
			})
		}
	}
}

func ShowCustomerWithOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	var customer entity.Customer

	err := models.ShowCustomerWithOrder(&customer, id)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	} else {
		if customer.Username == "" {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Customer not found",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"customer": customer,
			})
		}
	}
}

func UpdateCustomer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	var customer entity.CustomerResponse

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "cannot handle request",
		})
		return
	}

	customer.UpdateAt = time.Now()
	result := config.DB.Where("id = ? and is_active = ?", id, true).Updates(&customer)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "customer not found",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "data updated successfully",
		})
	}
}

func DeleteCustomer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	var customer entity.CustomerResponse

	result := config.DB.Model(&customer).Where("id = ? and is_active = ?", id, true).Update("is_active", false)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "customer not found",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "data deleted successfully",
		})
	}
}
