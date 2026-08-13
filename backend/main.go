package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() // デフォルトルーターを初期化
	// 引数に、パス、無名関数の順に指定する。
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run() // サーバーを起動(デフォルトポート：8080)
}
