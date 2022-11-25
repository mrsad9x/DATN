package controller

import (
	"DATN/service"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
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
		//c.JSONP(200, gin.H{
		//	"list": listProduct,
		//})
		c.HTML(200, "index.html", gin.H{
			"title":       "test",
			"listProduct": listProduct,
		})
	}
}

func (p ProductController) GetOneProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := p.pController.GetOneProduct(id)
	if err != nil {
		log.Println(err.Error())
	}
	c.JSONP(200, gin.H{
		"product": product,
	})

}

func (p ProductController) GetListProduct(c *gin.Context) {
	idDanhMuc, _ := strconv.Atoi(c.Param("id"))
	listProduct, err := p.pController.GetListProductWithCategories(idDanhMuc)
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
	nameSearch := c.Param("name")
	listProduct, err := p.pController.SearchProduct(nameSearch)
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
	idDanhMuc, _ := strconv.Atoi(c.PostForm("idDM"))
	tenSP := c.PostForm("tensp")
	giaBan, _ := strconv.ParseFloat(c.PostForm("giaban"), 64)
	giaNhap, _ := strconv.ParseFloat(c.PostForm("gianhap"), 64)
	soluong, _ := strconv.Atoi(c.PostForm("soluong"))
	mota := c.PostForm("mota")
	status := 1
	err := p.pController.CreateNewProduct(idDanhMuc, tenSP, giaBan, giaNhap, soluong, mota, status)
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
	id, _ := strconv.Atoi(c.PostForm("id"))
	idDanhMuc, _ := strconv.Atoi(c.PostForm("idDM"))
	tenSP := c.PostForm("tensp")
	giaBan, _ := strconv.ParseFloat(c.PostForm("giaban"), 64)
	giaNhap, _ := strconv.ParseFloat(c.PostForm("gianhap"), 64)
	soluong, _ := strconv.Atoi(c.PostForm("soluong"))
	mota := c.PostForm("mota")

	err := p.pController.AlterProduct(id, idDanhMuc, tenSP, giaBan, giaNhap, soluong, mota)
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
	id, _ := strconv.Atoi(c.PostForm("id"))
	err := p.pController.DeleteSoftProduct(id)
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
