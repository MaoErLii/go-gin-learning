package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"gin-demo/config"
)

// DB is ...
var DB *gorm.DB

// Init is 连接mysql数据库
func Init() {
	fmt.Println("mysql init")
	var err error

	DB, err = gorm.Open(mysql.Open(config.DSN), &gorm.Config{})

	if err != nil {
		fmt.Println("fail to connect database")
		panic("stop services")
	} else {
		fmt.Println("database connect success")
	}
}
