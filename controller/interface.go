package controller

import "github.com/gin-gonic/gin"

type IUserController interface {
	SetRouterUserController(router *gin.Engine) *gin.Engine
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
}
