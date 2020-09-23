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

// FilePath is 文件存储地址
var FilePath string

// 图片地址
var ImgUrl string

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT = os.Getenv("HTTP_PORT")
	DSN = os.Getenv("DSN")
	FilePath = os.Getenv("FILE_PATH")
	ImgUrl = os.Getenv("IMG_URL")
	log.Print(PORT)
}
