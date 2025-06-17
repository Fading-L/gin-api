package api

import (
	"github.com/Fading-L/gin-api/database"
	"github.com/Fading-L/gin-api/models"
	"github.com/gin-gonic/gin"
)


func Delete(c *gin.Context) {
	var id = c.Param("id")
	// 这里可以添加删除逻辑
	 database.Get().Delete(&models.Product{},id)

	c.JSON(200, gin.H{
		"message": "删除成功",
	})
}
