package controllers

import (
	"gin-demo/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UploadVideo is 视频上传方法
func UploadVideo(c *gin.Context) {
	file, _ := c.FormFile("video")

	if file == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "上传失败",
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

	e := c.SaveUploadedFile(file, config.FilePath+"/"+file.Filename)

	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": e.Error(),
		})
		return
	}

	var data uploadRes
	data.URL = config.ImgUrl + file.Filename

	c.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"meessage": "视频上传成功",
		"data":     data,
	})
}
