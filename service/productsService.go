package service

import (
	"DATN/model"
	"DATN/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 10*1024*1024)
	idDanhMuc, _ := strconv.Atoi(c.PostForm("idDM"))
	tenSP := c.PostForm("tensp")
	giaBan, _ := strconv.ParseFloat(c.PostForm("giaban"), 64)
	giaNhap, _ := strconv.ParseFloat(c.PostForm("gianhap"), 64)
	soluong, _ := strconv.Atoi(c.PostForm("soluong"))
	mota := c.PostForm("mota")
	imgRaw, err := c.FormFile("anh")
	if err != nil {
		return err
	}

	imgName := imgRaw.Filename

	fmt.Println(imgRaw)
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
