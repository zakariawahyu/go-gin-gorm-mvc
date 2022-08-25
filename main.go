package main

import (
	"github.com/zakariawahyu/go-gin-gorm-mvc/config"
	"github.com/zakariawahyu/go-gin-gorm-mvc/router"
)

func main() {
	config.DatabaseInit()
	router.Run()
}
