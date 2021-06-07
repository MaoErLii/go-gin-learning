package tables

import (
	"gin-demo/app/models"
	"gin-demo/database"
)

func initArticleTable() {
	hasTable := database.DB.Migrator().HasTable(&models.Article{})
	if hasTable {
		// database.DB.Migrator().DropTable(&models.Article{})
	}
	database.DB.Migrator().CreateTable(&models.Article{})
}
