package service

import (
	"DATN/model"
	"github.com/gin-gonic/gin"
)

type IUserService interface {
	Register(c *gin.Context) error
	Login(c *gin.Context) (string, error)
	CreateUser(c *gin.Context) error
	CheckRoles(token string) (int, error)
	ShowListUer() ([]model.User, error)
	AlterUser(c *gin.Context) error
}

type IProductService interface {
	GetAllProduct() ([]model.SanPham, error)
	GetOneProduct(c *gin.Context) (model.SanPham, error)
	GetListProductWithCategories(c *gin.Context) ([]model.SanPham, error)
	SearchProduct(c *gin.Context) ([]model.SanPham, error)
	CreateNewProduct(c *gin.Context) error
	AlterProduct(c *gin.Context) error
	DeleteSoftProduct(c *gin.Context) error
}

type IHomeService interface {
	Home() ([]model.SanPham, error)
}

type ICategoriesService interface {
}

type ICartService interface {
}

type IOrderService interface {
}

type IPromotionService interface {
}
