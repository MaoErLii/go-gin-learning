package main

import (
	"gin-demo/config"
	"gin-demo/routes"
)

func main() {

	// database.DB.Begin()

	// err := database.DB.Ping()

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// defer database.DB.Close()

	var router = routes.Init()

	router.Run(":" + config.PORT)
}
