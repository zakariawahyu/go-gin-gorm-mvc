package models

import (
	"github.com/zakariawahyu/go-gin-gorm-mvc/config"
	"github.com/zakariawahyu/go-gin-gorm-mvc/entity"
)

func GetAllProduct(product *[]entity.Product) (err error) {
	if err := config.DB.Find(&product).Error; err != nil {
		return err
	}
	return nil
}

func CreateProduct(product *entity.Product) (err error) {
	if err := config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func GetProductByID(product *entity.Product, id int) (err error) {
	if err := config.DB.Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}
