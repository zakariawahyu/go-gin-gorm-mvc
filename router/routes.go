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
	getRoutesOrderDetail(v1)
	getRouterPayment(v1)
	router.Run("localhost:8081")
}

func getRoutesCustomer(rg *gin.RouterGroup) {
	customer := rg.Group("/customer")
	customer.GET("/", controller.GetCustomers)
	customer.GET("/:id", controller.ShowCustomer)
	customer.GET("/detail-order", controller.GetCustomersWithOrder)
	customer.GET("/detail-order/:id", controller.ShowCustomerWithOrder)
	customer.POST("/", controller.CreateCustomers)
	customer.PUT("/:id", controller.UpdateCustomer)
	customer.DELETE("/:id", controller.DeleteCustomer)
}

func getRoutesProduct(rg *gin.RouterGroup) {
	product := rg.Group("/product")
	product.GET("/", controller.GetProduct)
	product.POST("/", controller.CreateProduct)
	product.GET("/:id", controller.ShowProduct)
	product.PUT("/:id", controller.UpdateProduct)
	product.DELETE("/:id", controller.DeleteProduct)
}

func getRoutesOrder(rg *gin.RouterGroup) {
	order := rg.Group("/order")
	order.GET("/", controller.GetOrder)
	order.POST("/", controller.CreateOrder)
	order.GET("/:id", controller.GetOrderById)
}

func getRoutesOrderDetail(rg *gin.RouterGroup) {
	orderDetail := rg.Group("/order-detail")
	orderDetail.GET("/", controller.GetOrderDetail)
}

func getRouterPayment(rg *gin.RouterGroup) {
	payment := rg.Group("/payment")
	payment.POST("/", controller.CreatePayment)
}
