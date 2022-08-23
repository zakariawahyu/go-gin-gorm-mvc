package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zakariawahyu/go-gin-gorm-mvc/controller"
)

func Run() {
	router := gin.Default()
	v1 := router.Group("/v1")
	getRoutes(v1)
	router.Run("localhost:8081")
}

func getRoutes(rg *gin.RouterGroup) {
	route := rg.Group("/customer")

	route.GET("/", controller.GetCustomers)
}
