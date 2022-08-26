package models

import (
	"github.com/zakariawahyu/go-gin-gorm-mvc/config"
	"github.com/zakariawahyu/go-gin-gorm-mvc/entity"
)

func GetAllOrderDetail(orderDetail *[]entity.OrderDetail) (err error) {
	if err = config.DB.Preload("Product").Preload("Order.Customer").Find(orderDetail).Error; err != nil {
		return err
	}
	return nil
}

func CreateOrderDetail(orderDetail *entity.OrderDetail) (err error) {
	if err := config.DB.Create(orderDetail).Error; err != nil {
		return err
	}
	return nil
}
