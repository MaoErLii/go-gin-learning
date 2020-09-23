package controllers

import (
	"fmt"
	"net/http"
	"path"

	"gin-demo/config"

	"github.com/gin-gonic/gin"
)

type uploadRes struct {
	URL string `json:"url" `
}

// UploadImage is ...单图片上传，接收上传图片，并存入指定文件夹中
func UploadImage(c *gin.Context) {
	file, _ := c.FormFile("image")

	if file == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "上传失败",
		})
		return
	}

	fileSize := file.Size
	fmt.Println("上传图片大小", fileSize)
	fileExt := path.Ext(file.Filename)
	fmt.Println("上传文件后缀", fileExt)
	if fileExt != ".jpg" && fileExt != ".png" && fileExt != ".bmp" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "上传图片格式错误",
		})
		return
	}

	if !IsDirExist(config.FilePath) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "上传失败",
		})
		return
	}

	err := c.SaveUploadedFile(file, config.FilePath+"/"+file.Filename)

	if err != nil {
		// fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	var data uploadRes
	data.URL = config.ImgUrl + file.Filename

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "上传成功",
		"data":    data,
	})
}
