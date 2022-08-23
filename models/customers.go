package models

import (
	"github.com/zakariawahyu/go-gin-gorm-mvc/config"
	"github.com/zakariawahyu/go-gin-gorm-mvc/entity"
)

func GetAllCustomers(customer *[]entity.Customer) (err error) {
	if err = config.DB.Find(customer).Error; err != nil {
		return err
	}
	return nil
}
