package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"go-api/model"
)

func OpenDb() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(goDb:3306)/memodb?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	return db
}

func DbInit() {
	db := OpenDb()
	db.AutoMigrate(&model.Todo{})
}
