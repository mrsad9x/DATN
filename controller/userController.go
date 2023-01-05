package controller

import (
	"DATN/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

type UserController struct {
	UController service.IUserService
}

func NewUserController(userControl service.IUserService) IUserController {
	return UserController{
		UController: userControl,
	}
}

func (u UserController) SetRouterUserController(router *gin.Engine) *gin.Engine {
	router.POST("/login", u.Login)
	router.POST("/register", u.Register)
	r := router.Group("/admin")
	{
		r.POST("/createuser", u.CreateUser)
		r.GET("/listuser", u.ListUser)
	}
	return router
}

func (u UserController) Login(c *gin.Context) {
	username := c.PostForm("user")
	pass := c.PostForm("pass")
	token, err := u.UController.Login(username, pass)

	if err != nil {
		fmt.Println(err)
		c.JSONP(400, gin.H{
			"messeage": "login fail",
		})
	} else {
		http.SetCookie(c.Writer, &http.Cookie{
			Name:    "token",
			Value:   token,
			Expires: time.Now().Add(2 * time.Minute),
		})
		c.JSONP(http.StatusOK, gin.H{
			"user":  username,
			"token": token,
		})
	}
}

func (u UserController) Register(c *gin.Context) {
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
		c.JSONP(http.StatusOK, gin.H{
			"message": "register success!",
		})
	}
}

func (u UserController) CreateUser(c *gin.Context) {
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
		c.JSONP(http.StatusOK, gin.H{
			"message": "register success!",
		})
	}
}

func (u UserController) UpdateInfo(c *gin.Context) {

}

func (u UserController) ListUser(c *gin.Context) {
	role, err := u.checkUser(c)
	if err != nil {
		c.JSONP(http.StatusNetworkAuthenticationRequired, "")
		return
	}
	if role == 3 {
		c.JSONP(http.StatusUnauthorized, gin.H{
			"role": role,
		})
		return
	}
	listUser, err := u.UController.ShowListUer()
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"message": "get list user fail",
		})
		return
	}
	c.JSONP(http.StatusOK, gin.H{
		"list user": listUser,
	})
}

func (u UserController) checkUser(c *gin.Context) (int, error) {
	cookie, err := c.Request.Cookie("token")
	if err != nil {
		return 0, err
	}
	return u.UController.CheckRoles(cookie.Value)

}
