package models

import (
	"github.com/zakariawahyu/go-gin-gorm-mvc/config"
	"github.com/zakariawahyu/go-gin-gorm-mvc/entity"
)

func GetAllProduct(product *[]entity.Product) (err error) {
	if err := config.DB.Where("is_active = ?", true).Find(&product).Error; err != nil {
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

func ShowProduct(product *entity.Product, id int) (err error) {
	if err := config.DB.Where("id = ? and is_active = ?", id, true).First(product).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProduct(product *entity.Product, id int) (err error) {
	if err := config.DB.Where("id = ?", id).Updates(product).Error; err != nil {
		return err
	}
	return nil
}

func DeleteProduct(product *entity.Product, id int) (err error) {
	if err := config.DB.Model(product).Where("id = ?", id).UpdateColumn("is_active", false).Error; err != nil {
		return err
	}
	return nil
}
