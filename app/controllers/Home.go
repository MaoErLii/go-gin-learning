package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Home is ...
func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title": "hello world",
	})
}
