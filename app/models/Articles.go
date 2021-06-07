package models

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	Article  Model `gorm:"embedded;embeddedPrefix:article_"`
	FilePath string
}

// 新增文章
func (article *Article) AddArticle(db *gorm.DB) (err error) {
	result := db.Create(&article)

	return result.Error
}

// 更新文章
func UpdateArticle(filePath string, db *gorm.DB) (err error) {
	articleModel := Model{UpdateAt: time.Now()}
	result := db.Model(&Article{}).Where("file_path = ?", filePath).Updates(Article{FilePath: "filePath", Article: articleModel})

	return result.Error
}
