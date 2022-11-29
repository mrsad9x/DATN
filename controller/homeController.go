package controller

import (
	"DATN/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HomeController struct {
	hController service.IHomeService
}

func NewHomeController(homeController service.IHomeService) IHomeController {
	return HomeController{hController: homeController}
}

func (h HomeController) SetRouterHomeController(router *gin.Engine) *gin.Engine {
	router.GET("/", h.Home)
	router.GET("/home", h.Home)
	return router
}

func (h HomeController) Home(c *gin.Context) {
	listPro, err := h.hController.Home()
	if err != nil {
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"listProduct": listPro,
		})
	}

}
