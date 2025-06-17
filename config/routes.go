 package config

 import (
  "github.com/gin-gonic/gin"
  "github.com/Fading-L/gin-api/api"
)
func Routes(r *gin.Engine)  {
  
  // 这里可以添加路由规则
  //  r := gin.Default()
 r.GET("/ping",  api.Ping )
 r.GET("/index", api.Index)
 r.POST("/create", api.Create)
 r.POST("/delete/:id", api.Delete)
 r.POST("/update/:id", api.Update)
 
}