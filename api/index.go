package api

import (
	"net/http"

	"github.com/Fading-L/gin-api/database"
	"github.com/Fading-L/gin-api/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var products []models.Product
	// 连接数据库
	 database.Get().Find(&products)

	c.JSON(http.StatusOK, gin.H{"list": products})
	
	
}
