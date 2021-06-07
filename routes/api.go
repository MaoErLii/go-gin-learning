package routes

import (
	"fmt"

	"gin-demo/app/controllers"
	"gin-demo/app/middleware"

	"github.com/gin-gonic/gin"
)

// Init is ...
func Init() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CorsHandler())

	fmt.Println("go router")

	v1 := r.Group("/")
	{
		v1.GET("/home", controllers.Home)
	}

	// v2 := r.Group("/user")
	// {
	// 	v2.GET("/info", controllers.UserInfo)
	// 	v2.POST("/add", controllers.AddUser)
	// 	v2.DELETE("/delete", controllers.DeleteUser)
	// 	v2.PUT("/update", controllers.UpdateUser)
	// }

	v3 := r.Group("/article")
	{
		v3.POST("/latest")
		v3.POST("/detail")
		v3.POST("/list")
		v3.POST("/upload", controllers.UploadArticle)
	}

	v4 := r.Group("image")
	{
		v4.POST("/upload", controllers.UploadImage)
	}

	v5 := r.Group("video")
	{
		v5.POST("/upload", controllers.UploadVideo)
	}

	return r
}
