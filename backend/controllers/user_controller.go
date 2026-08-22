package controllers

import (
	"backend/dto"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IUserController interface {
	FindById(ctx *gin.Context)
	Create(ctx *gin.Context)
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

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (c *UserController) Create(ctx *gin.Context) {
	var request dto.CreateUserInput
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.service.Create(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"user": user})
}
