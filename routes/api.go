package routes

import (
	"fmt"
	"gin-demo/app/controllers"

	"github.com/gin-gonic/gin"
)

// Init is ...
func Init() *gin.Engine {
	r := gin.Default()

	fmt.Println("go router")

	v1 := r.Group("/")
	{
		v1.GET("/home", controllers.Home)
	}

	v2 := r.Group("/user")
	{
		v2.GET("/info", controllers.UserInfo)
		v2.POST("/add", controllers.AddUser)
		v2.DELETE("/delete", controllers.DeleteUser)
	}

	return r
}
