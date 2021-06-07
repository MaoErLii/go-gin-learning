package controllers

import (
	"gin-demo/app/models"
	"net/http"
	"os"
	"time"

	"gin-demo/database"

	"github.com/gin-gonic/gin"
)

// 文章上传
func UploadArticle(context *gin.Context) {
	file, _ := context.FormFile("file")
	dir, _ := os.Getwd()
	articleFilePath := dir + "\\files\\articles\\" + file.Filename
	// 检测路径真实性
	if !IsDirExist(dir) {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "存储路径不存在",
		})
		return
	}
	// 不存在文件，存入files文件中
	if !IsDirExist(articleFilePath) {
		err := context.SaveUploadedFile(file, articleFilePath)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}
		var articleModel = models.Model{CreateAt: time.Now()}
		var article = models.Article{FilePath: articleFilePath, Article: articleModel}
		// 插入表中
		article.AddArticle(database.DB)
	} else {
		// 存在文件更新文章
		models.UpdateArticle(articleFilePath, database.DB)
	}

}

// 获取文章列表
func GetArticles(ctx *gin.Context) {

}
