package models

import (
	"github.com/zakariawahyu/go-gin-gorm-mvc/config"
	"github.com/zakariawahyu/go-gin-gorm-mvc/entity"
)

func CreatePayment(payment *entity.Payment) error {
	if err := config.DB.Create(payment).Error; err != nil {
		return err
	}
	return nil
}
