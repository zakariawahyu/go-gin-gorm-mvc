package models

import (
	"github.com/zakariawahyu/go-gin-gorm-mvc/config"
	"github.com/zakariawahyu/go-gin-gorm-mvc/entity"
)

func GetAllCustomers(customer *[]entity.Customer) (err error) {
	if err = config.DB.Preload("Order.OrderDetail.Product").Find(customer).Error; err != nil {
		return err
	}
	return nil
}

func CreateCustomers(customer *entity.Customer) (err error) {
	if err := config.DB.Create(customer).Error; err != nil {
		return err
	}
	return nil
}
