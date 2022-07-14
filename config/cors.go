package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:28000",
		},
		// アクセス許可するHTTPメソッド
		AllowMethods: []string{
			"POST",
			"GET",
		},
		// 許可するHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Content-Type",
		},
	}))
}
