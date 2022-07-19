// モデル
package model

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Todo struct {
	ID        int            `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
	Memo      string         `json:"memo"`
}

var openDb *gorm.DB

func init() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(goDb:3306)/memodb?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	openDb = db
}
func GetAll() *[]Todo {
	var todos *[]Todo
	openDb.Find(&todos)
	return todos
}

func CreateTodo(memo string) *[]Todo {
	openDb.Create(&Todo{Memo: memo})

	var todo *[]Todo
	openDb.Last(&todo)

	return todo
}

func Delete(id int) *gorm.DB {
	result := openDb.Delete(&Todo{ID: id})
	return result
}

func Update(id int, memo string) *gorm.DB {
	fmt.Println(memo)
	result := openDb.Model(&Todo{ID: id}).Updates(Todo{Memo: memo})
	return result
}
