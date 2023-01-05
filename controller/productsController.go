package controller

import (
	"DATN/service"
	"github.com/gin-gonic/gin"
	"log"
)

type ProductController struct {
	pController service.IProductService
}

func NewProductController(productController service.IProductService) IProductController {
	return ProductController{
		pController: productController,
	}
}

func (p ProductController) SetRouterSanPhamController(router *gin.Engine) *gin.Engine {

	router.GET("/getall", p.GetAllProduct)
	router.GET("/product/:id", p.GetOneProduct)
	router.GET("/categories/:id", p.GetListProduct)
	router.GET("/search/:name", p.SearchProduct)

	r := router.Group("/admin")
	{
		r.POST("/createproduct", p.CreateNewProduct)
		r.PUT("/alterproduct", p.AlterProduct)
		r.PUT("/deletesoftprod", p.DeleteSoftProduct)
	}
	return router
}

func (p ProductController) GetAllProduct(c *gin.Context) {
	listProduct, err := p.pController.GetAllProduct()
	if err != nil {
		c.JSONP(400, gin.H{
			"title": err,
		})
	} else {
		c.JSONP(200, gin.H{
			"listProduct": listProduct,
		})
		//c.HTML(200, "index.html", gin.H{
		//	"title":       "test",
		//	"listProduct": listProduct,
		//})
	}
}

func (p ProductController) GetOneProduct(c *gin.Context) {

	product, err := p.pController.GetOneProduct(c)
	if err != nil {
		log.Println(err.Error())
	}
	c.JSONP(200, gin.H{
		"product": product,
	})

}

func (p ProductController) GetListProduct(c *gin.Context) {
	listProduct, err := p.pController.GetListProductWithCategories(c)
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

func (p ProductController) SearchProduct(c *gin.Context) {

	listProduct, err := p.pController.SearchProduct(c)
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

func (p ProductController) CreateNewProduct(c *gin.Context) {
	err := p.pController.CreateNewProduct(c)
	if err != nil {
		c.JSONP(400, gin.H{
			"err": err,
		})
	} else {
		c.JSONP(200, gin.H{
			"message": "insert success",
		})
	}
}

func (p ProductController) AlterProduct(c *gin.Context) {
	err := p.pController.AlterProduct(c)
	if err != nil {
		c.JSONP(400, gin.H{
			"err": err,
		})
	} else {
		c.JSONP(200, gin.H{
			"msg": "thanh cong",
		})
	}
}

func (p ProductController) DeleteSoftProduct(c *gin.Context) {
	err := p.pController.DeleteSoftProduct(c)
	if err != nil {
		c.JSONP(400, gin.H{
			"err": err,
		})
	} else {
		c.JSONP(200, gin.H{
			"msg": "thanh cong",
		})
	}
}
