package controller

import (
	"DATN/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UController service.IUserService
}

func NewUserController(userControl service.IUserService) INguoiDungController {
	return new(UserController)
}

func (u UserController) SetRouterUserController(router *gin.Engine) *gin.Engine {
	router.GET("/login", u.DangNhap)
	r := router.Group("/admin")
	{
		r.GET("/")
	}
	return router
}

func (u UserController) DangNhap(c *gin.Context) {

}
