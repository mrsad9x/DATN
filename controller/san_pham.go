package controller

import (
	"DATN/service"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type ProductController struct {
	PController service.IProductService
}

func NewProductController(productController service.IProductService) ISanPhamController {
	return ProductController{
		PController: productController,
	}
}

func (p ProductController) SetRouterSanPhamController(router *gin.Engine) *gin.Engine {
	router.GET("/getall", p.GetAllProduct)
	router.GET("/getproduct/:id", p.GetOneProduct)
	return router
}

func (p ProductController) GetAllProduct(c *gin.Context) {
	listProduct, err := p.PController.GetAllProduct()
	if err != nil {
		c.JSONP(400, gin.H{
			"error": err,
		})
	} else {
		c.JSONP(200, gin.H{
			"list": listProduct,
		})
	}
}

func (p ProductController) GetOneProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := p.PController.GetOneProduct(id)
	if err != nil {
		log.Println(err.Error())
	}
	c.JSONP(200, gin.H{
		"product": product,
	})

}

func (p ProductController) GetListProduct(c *gin.Context) {

}
