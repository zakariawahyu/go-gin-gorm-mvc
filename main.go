package main

import (
	"github.com/zakariawahyu/go-gin-gorm-mvc/config"
	"github.com/zakariawahyu/go-gin-gorm-mvc/entity"
	"github.com/zakariawahyu/go-gin-gorm-mvc/router"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open(mysql.Open(config.DbURL(config.BuildDBConfig())), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	defer config.DB.AutoMigrate(&entity.Customer{})

	router.Run()
}
