package main

import (
 "github.com/gin-gonic/gin"
 "github.com/Fading-L/gin-api/config"
 "github.com/Fading-L/gin-api/database"
)

func main() {
  r := gin.Default()
  	database.Setup()
	config.Routes(r)
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}