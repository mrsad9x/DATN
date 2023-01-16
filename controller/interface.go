package controller

import "github.com/gin-gonic/gin"

type IUserController interface {
	SetRouterUserController(router *gin.Engine) *gin.Engine
	CheckUser(c *gin.Context) (int, error)
}

type IProductController interface {
	SetRouterSanPhamController(router *gin.Engine) *gin.Engine
}

type IHomeController interface {
	SetRouterHomeController(router *gin.Engine) *gin.Engine
}

type ICategoriesController interface {
}

type ICartController interface {
}

type IOrderController interface {
}

type IPromotionController interface {
	SetRouterPromotionController(router *gin.Engine) *gin.Engine
}
