package main

import (
	"fmt"
	"gin-demo/config"
	"gin-demo/database"
	"gin-demo/database/tables"
	"gin-demo/routes"
)

func main() {
	fmt.Println("main running")

	config.Init()

	err := database.Init()
	if err != nil {
		fmt.Println("fail to connect database")
		panic("stop services")
	} else {
		fmt.Println("database connect success")

		fmt.Println("数据库名", database.DB.Migrator().CurrentDatabase())

		// 初始化表
		tables.Init()

		var router = routes.Init()

		// router.Use(middleware.CorsHandler())

		router.Run(":" + config.PORT)

	}

	// database.DB.Begin()

	// err := database.DB.Ping()

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// defer database.DB.Close()
}
