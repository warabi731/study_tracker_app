package controllers

import (
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IUserController interface {
	FindById(ctx *gin.Context)
}

type UserController struct {
	service services.IUserService
}

func NewUserController(service services.IUserService) IUserController {
	return &UserController{service: service}
}

func (c *UserController) FindById(ctx *gin.Context) {
	// [TODO]id はCookie情報から取得する。一旦はid=1で固定で取得する
	var userId uint = 1

	user, err := c.service.FindById(userId)
	if err != nil {
		if err.Error() == "user not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "user": user })
}
