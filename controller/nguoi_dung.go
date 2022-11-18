package controller

import (
	"DATN/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UController service.IUserService
}

func NewUserController(userControl service.IUserService) INguoiDungController {
	return UserController{
		UController: userControl,
	}
}

func (u UserController) SetRouterUserController(router *gin.Engine) *gin.Engine {
	router.POST("/login", u.DangNhap)
	router.POST("/register", u.DangKy)
	r := router.Group("/admin")
	{
		r.GET("/")
	}
	return router
}

func (u UserController) DangNhap(c *gin.Context) {
	username := c.PostForm("user")
	pass := c.PostForm("pass")
	err := u.UController.Login(username, pass)
	if err != nil {
		fmt.Println(err)
	}
	c.JSONP(200, gin.H{
		"message": "login success!",
	})
}

func (u UserController) DangKy(c *gin.Context) {

}
