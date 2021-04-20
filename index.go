package main

import (
	"fmt"
	"gin-demo/app/middleware"
	"gin-demo/config"
	"gin-demo/database"
	"gin-demo/routes"
)

func main() {
	fmt.Println("main running")

	config.Init()

	database.Init()

	// database.DB.Begin()

	// err := database.DB.Ping()

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// defer database.DB.Close()

	var router = routes.Init()

	router.Use(middleware.CorsHandler())

	router.Run(":" + config.PORT)
}
