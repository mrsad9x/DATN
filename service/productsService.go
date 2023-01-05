package service

import (
	"DATN/model"
	"DATN/repository"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ProductService struct {
	proService repository.IProductDB
}

func NewProducService(repo repository.IProductDB) IProductService {
	return ProductService{proService: repo}
}

func (p ProductService) GetAllProduct() ([]model.SanPham, error) {
	return p.proService.GetAllProduct()
}

func (p ProductService) GetOneProduct(c *gin.Context) (model.SanPham, error) {
	id, _ := strconv.Atoi(c.Param("id"))
	return p.proService.GetOneProduct(id)
}

func (p ProductService) GetListProductWithCategories(c *gin.Context) ([]model.SanPham, error) {
	idDanhMuc, _ := strconv.Atoi(c.Param("id"))
	return p.proService.GetListProductWithCategories(idDanhMuc)
}

func (p ProductService) SearchProduct(c *gin.Context) ([]model.SanPham, error) {
	nameSearch := c.Param("name")
	return p.proService.SearchProduct(nameSearch)
}

func (p ProductService) CreateNewProduct(c *gin.Context) error {
	idDanhMuc, _ := strconv.Atoi(c.PostForm("idDM"))
	tenSP := c.PostForm("tensp")
	giaBan, _ := strconv.ParseFloat(c.PostForm("giaban"), 64)
	giaNhap, _ := strconv.ParseFloat(c.PostForm("gianhap"), 64)
	soluong, _ := strconv.Atoi(c.PostForm("soluong"))
	mota := c.PostForm("mota")
	status := 1
	return p.proService.CreateNewProduct(idDanhMuc, tenSP, giaBan, giaNhap, soluong, mota, status)
}

func (p ProductService) AlterProduct(c *gin.Context) error {
	id, _ := strconv.Atoi(c.PostForm("id"))
	idDanhMuc, _ := strconv.Atoi(c.PostForm("idDM"))
	tenSP := c.PostForm("tensp")
	giaBan, _ := strconv.ParseFloat(c.PostForm("giaban"), 64)
	giaNhap, _ := strconv.ParseFloat(c.PostForm("gianhap"), 64)
	soluong, _ := strconv.Atoi(c.PostForm("soluong"))
	mota := c.PostForm("mota")
	return p.proService.AlterProduct(id, idDanhMuc, tenSP, giaBan, giaNhap, soluong, mota)
}

func (p ProductService) DeleteSoftProduct(c *gin.Context) error {
	id, _ := strconv.Atoi(c.PostForm("id"))
	status := 0
	return p.proService.DeleteSoftProduct(id, status)
}
