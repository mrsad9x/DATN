package controller

import (
	"DATN/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
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
		r.POST("/createuser", u.TaoNguoiDung)
	}
	return router
}

func (u UserController) DangNhap(c *gin.Context) {
	username := c.PostForm("user")
	pass := c.PostForm("pass")
	err := u.UController.Login(username, pass)
	if err != nil {
		fmt.Println(err)
		c.JSONP(400, gin.H{
			"messeage": "login fail",
		})
	} else {
		c.JSONP(200, gin.H{
			"message": "login success!",
		})
	}
}

func (u UserController) DangKy(c *gin.Context) {
	ten := c.PostForm("name")
	taiKhoan := c.PostForm("username")
	matKhau := c.PostForm("password")
	sdt := c.PostForm("phone")
	email := c.PostForm("email")
	diaChi := ""
	chiSo := 1
	status := 1
	role := 3
	err := u.UController.Register(ten, taiKhoan, matKhau, sdt, email, diaChi, status, role, chiSo)
	if err != nil {
		log.Println(err.Error())
		c.JSONP(400, gin.H{
			"message": "register fail!",
		})
	} else {
		c.JSONP(200, gin.H{
			"message": "register success!",
		})
	}
}

func (u UserController) TaoNguoiDung(c *gin.Context) {
	ten := c.PostForm("name")
	taiKhoan := c.PostForm("username")
	matKhau := c.PostForm("password")
	sdt := c.PostForm("phone")
	email := c.PostForm("email")
	diaChi := c.PostForm("place")
	chiSo := 1
	status := 1
	role, _ := strconv.Atoi(c.PostForm("role"))
	err := u.UController.Register(ten, taiKhoan, matKhau, sdt, email, diaChi, status, role, chiSo)
	if err != nil {
		log.Println(err.Error())
		c.JSONP(400, gin.H{
			"message": "register fail!",
		})
	} else {
		c.JSONP(200, gin.H{
			"message": "register success!",
		})
	}
}

func (u UserController) SuaNguoiDung(c *gin.Context) {

}
