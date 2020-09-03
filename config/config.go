package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// PORT is ... 服务运行端口
var PORT string

// DSN is ... 用户连接数据库信息
// 格式参考 gorm 文档数据库连接教程
var DSN string

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT = os.Getenv("HTTP_PORT")
	DSN = os.Getenv("DSN")
	log.Print(PORT)
}
