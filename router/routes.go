package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zakariawahyu/go-gin-gorm-mvc/controller"
)

func Run() {
	router := gin.Default()
	v1 := router.Group("/v1")
	getRoutesCustomer(v1)
	getRoutesProduct(v1)
	getRoutesOrder(v1)
	router.Run("localhost:8081")
}

func getRoutesCustomer(rg *gin.RouterGroup) {
	customer := rg.Group("/customer")
	customer.GET("/", controller.GetCustomers)
}

func getRoutesProduct(rg *gin.RouterGroup) {
	product := rg.Group("/product")
	product.GET("/", controller.GetProduct)
}

func getRoutesOrder(rg *gin.RouterGroup) {
	order := rg.Group("/order")
	order.GET("/", controller.GetOrder)
}
