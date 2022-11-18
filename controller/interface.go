package controller

import "github.com/gin-gonic/gin"

type INguoiDungController interface {
	SetRouterUserController(router *gin.Engine) *gin.Engine
}
