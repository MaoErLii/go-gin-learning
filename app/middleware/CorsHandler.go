package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 跨域处理中间件
func CorsHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method

		fmt.Println(context)

		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Origin", "*") // 设置允许访问所有域
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		context.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma,token,openid,opentoken")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
		context.Header("Access-Control-Max-Age", "172800")
		context.Header("Access-Control-Allow-Credentials", "false")
		context.Set("content-type", "application/json") // 设置返回格式是json

		if method == "OPTIONS" {
			context.JSON(http.StatusOK, gin.H{"code": "ok"})
		}

		context.Next()

		// ————————————————
		// 版权声明：本文为CSDN博主「伊宇紫」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
		// 原文链接：https://blog.csdn.net/LDY1016/article/details/106691752
	}
}
