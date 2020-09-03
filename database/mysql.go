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
func init() {
	fmt.Println("mysql init")
	var err error

	DB, err = gorm.Open(mysql.Open(config.DSN), &gorm.Config{})

	if err != nil {
		fmt.Println("mysql conect error v%", err)
	}

	if err != nil {
		panic("连接失败")
	} else {
		fmt.Println("数据库连接成功")
	}
}
