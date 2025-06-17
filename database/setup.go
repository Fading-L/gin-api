package database

import (
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
  "github.com/Fading-L/gin-api/models"
)



var gdb *gorm.DB // 全局变量 保存数据库连接

func Setup() {
  db, err := gorm.Open(sqlite.Open("./data/test.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  // 迁移 schema
  db.AutoMigrate(&models.Product{})
  gdb = db
}

func Get () *gorm.DB {
	if gdb == nil {
		panic("数据库未初始化")
	}
	return gdb
}