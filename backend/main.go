package main

import (
	"backend/controllers"
	"backend/repositories"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() // デフォルトルーターを初期化

	user_controller := CreateUserController()

	// 引数に、パス、無名関数の順に指定する。
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/users", user_controller.FindById)
	router.POST("/users", user_controller.Create)
	router.PUT("/users", user_controller.Update)

	router.Run() // サーバーを起動(デフォルトポート：8080)
}

func CreateUserController() controllers.IUserController {
	r := repositories.NewUserRepository()
	s := services.NewUserService(r)
	return controllers.NewUserController(s)
}
