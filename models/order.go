package models

import (
	"github.com/zakariawahyu/go-gin-gorm-mvc/config"
	"github.com/zakariawahyu/go-gin-gorm-mvc/entity"
)

func GetAllOrder(order *[]entity.Order) (err error) {
	if err = config.DB.Preload("Customer").Find(order).Error; err != nil {
		return err
	}
	return nil
}
