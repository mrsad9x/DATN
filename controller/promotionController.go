package controller

import (
	"DATN/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PromotionController struct {
	promoController service.IPromotionService
	uController     IUserController
}

func NewPromotionController(serv service.IPromotionService, user IUserController) IPromotionController {
	return PromotionController{promoController: serv, uController: user}
}

func (p PromotionController) SetRouterPromotionController(router *gin.Engine) *gin.Engine {

	r := router.Group("/admin")
	{
		r.POST("/createpromotion", p.CreatePromotion)
		r.PUT("/alterpromotion", p.AlterPromotion)
		r.POST("/deletepromotion", p.DeletePromotion)
	}
	return router
}

func (p PromotionController) CreatePromotion(c *gin.Context) {
	role, err := p.uController.CheckUser(c)
	if err != nil {
		return
	}
	if role != 1 && role != 2 {
		c.JSONP(http.StatusUnauthorized, gin.H{
			"role": role,
		})
		return
	}
	err = p.promoController.CreatePromotion(c)
	if err != nil {
		return
	}

}

func (p PromotionController) AlterPromotion(c *gin.Context) {

}

func (p PromotionController) DeletePromotion(c *gin.Context) {

}
