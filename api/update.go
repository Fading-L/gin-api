package api

import (
	"github.com/Fading-L/gin-api/api"
	"github.com/Fading-L/gin-api/database"
	"github.com/Fading-L/gin-api/models"
	"github.com/gin-gonic/gin"
)


func Update(c *gin.Context ){
	var params Params
	id := c.Param("id")
	if err := c.BindJSON(&params);err != nil{
		c.JSON(400,gin.H{"error":err.Error()})
		return
	}

	
	database.Get().Model(&models.Product{}).Where("id = ?", id).
	Updates(map[string]any{
		"name":params.Name,
		"price":params.Price})
		c.JSON(200,gin.H{"message":"update success"})


}