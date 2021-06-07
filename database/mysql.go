package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"gin-demo/config"
	// "gin-demo/app/models"
)

// DB is ...
var DB *gorm.DB

// Init is 连接mysql数据库
func Init() (err error) {
	fmt.Println("mysql init")
	// var err error

	DB, err = gorm.Open(mysql.Open(config.DSN), &gorm.Config{
		// 注意 AutoMigrate 会自动创建数据库外键约束，在初始化时禁用此功能
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	return err

	// if err != nil {
	// 	fmt.Println("fail to connect database")
	// 	panic("stop services")
	// } else {
	// 	fmt.Println("database connect success")

	// 	fmt.Println("数据库名", DB.Migrator().CurrentDatabase())
	// }
}
