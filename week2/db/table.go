package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	//_ = g.Db.AutoMigrate(&Product{})
	//db.Create(&Product{Price: 123,Code: "hello"})

	return database
}
