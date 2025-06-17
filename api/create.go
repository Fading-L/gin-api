package api

import (
	"github.com/Fading-L/gin-api/database"
	"github.com/Fading-L/gin-api/models"
	"github.com/gin-gonic/gin"
)

// CreateParams 创建商品的参数结构体
type Params struct {
	Name  string `json:"name"`
	Price uint   `json:"price"`
}

func Create(c *gin.Context) {
	var params Params
	if err := c.BindJSON(&params); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	product := models.Product{
		Name:  params.Name,
		Price: params.Price,
	}

	database.Get().Create(&product)

	c.JSON(200, gin.H{
		"product": product,
	})
}
