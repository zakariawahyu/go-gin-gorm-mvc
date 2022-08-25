package config

import (
	"fmt"
	"github.com/zakariawahyu/go-gin-gorm-mvc/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func buildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "masukdb",
		DBName:   "golang_gin_gorm_mvc",
	}

	return &dbConfig
}

func dbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

func DatabaseInit() {
	var err error
	DB, err = gorm.Open(mysql.Open(dbURL(buildDBConfig())), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	defer DB.AutoMigrate(
		&entity.Customer{},
		&entity.Product{},
		&entity.Order{},
		&entity.OrderDetail{},
		//&entity.Payment{}
	)
}
