package models

import (
	"github.com/zakariawahyu/go-gin-gorm-mvc/config"
	"github.com/zakariawahyu/go-gin-gorm-mvc/entity"
)

func GetAllCustomers(customer *[]entity.CustomerWithoutOrder) (err error) {
	if err = config.DB.Where("is_active = ?", true).Find(customer).Error; err != nil {
		return err
	}
	return nil
}

func GetAllCustomersWithOrder(customer *[]entity.Customer) (err error) {
	if err = config.DB.Where("is_active = ?", true).Preload("Order.OrderDetail.Product").Find(customer).Error; err != nil {
		return err
	}
	return nil
}

func CreateCustomers(customer *entity.CustomerResponse) (err error) {
	if err := config.DB.Create(customer).Error; err != nil {
		return err
	}
	return nil
}

func ShowCustomer(customer *entity.CustomerWithoutOrder, id int) (err error) {
	if err := config.DB.Where("id = ? and is_active = ?", id, true).Find(customer).Error; err != nil {
		return err
	}
	return nil
}

func ShowCustomerWithOrder(customer *entity.Customer, id int) (err error) {
	if err := config.DB.Where("id = ? and is_active = ?", id, true).Preload("Order.OrderDetail.Product").Find(customer).Error; err != nil {
		return err
	}
	return nil
}
