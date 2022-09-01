package models

import (
	"github.com/zakariawahyu/go-gin-gorm-mvc/config"
	"github.com/zakariawahyu/go-gin-gorm-mvc/entity"
)

func GetAllOrder(order *[]entity.OrderResponse) (err error) {
	if err = config.DB.Preload("Customer").Preload("OrderDetail.Product").Find(order).Error; err != nil {
		return err
	}
	return nil
}

func CreateOrder(order *entity.Order) (err error) {
	if err := config.DB.Create(order).Error; err != nil {
		return err
	}
	return nil
}

func GetOrderById(order *entity.OrderResponse, id int) error {
	if err := config.DB.Where("id = ?", id).Preload("OrderDetail.Product").Preload("Customer").First(order).Error; err != nil {
		return err
	}
	return nil
}
