package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-api/model"
)

var postMemo model.Todo

func TodoAll(c *gin.Context) {

	c.JSONP(200, gin.H{
		"data": model.GetAll(),
	})
}

func CreateTodo(c *gin.Context) {

	if err := c.ShouldBindJSON(&postMemo); err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := model.CreateTodo(postMemo.Memo)

	c.JSONP(200, gin.H{"data": result})
}

func DeleteTodo(c *gin.Context) {
	if err := c.ShouldBindJSON(&postMemo); err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	model.Delete(postMemo.ID)

	c.JSONP(200, gin.H{"data": "削除しました"})
}

func UpdateTodo(c *gin.Context) {
	if err := c.ShouldBindJSON(&postMemo); err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	model.Update(postMemo.ID, postMemo.Memo)
	c.JSONP(200, gin.H{"data": "更新しました"})
}
