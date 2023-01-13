package service

import (
	"DATN/model"
	"DATN/repository"
	"DATN/repository/s3"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type ProductService struct {
	proService repository.IProductDB
	upload     s3.IS3Repo
}

func NewProducService(repo repository.IProductDB, s3store *s3.IS3Repo) IProductService {
	return ProductService{proService: repo, upload: *s3store}
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
	imgRaw, handler, err := c.Request.FormFile("anh")
	if err != nil {
		return err
	}
	imgName := strings.Split(handler.Filename, ".")[0]
	fileType := strings.Split(handler.Header.Get("Content-Type"), "/")[1]
	imgName = fmt.Sprintf("%s.%s", imgName, fileType)
	tempFile, err := os.Create("temp-img/" + imgName)
	if err != nil {
		return err
	}

	fileBytes, err := ioutil.ReadAll(imgRaw)
	if err != nil {
		fmt.Println(err)
	}

	tempFile.Write(fileBytes)
	tempFile.Close()
	status := 1
	err = p.proService.CreateNewProduct(idDanhMuc, tenSP, giaBan, giaNhap, soluong, mota, status, imgName)
	if err != nil {
		return err
	}
	fileUpload, err := os.Open(fmt.Sprintf("temp-img/%s", imgName))

	pathS3 := "image/" + imgName
	err = p.upload.PutObject(pathS3, fileUpload)
	fileUpload.Close()
	err = os.Remove(fmt.Sprintf("temp-img/%s", imgName))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (p ProductService) AlterProduct(c *gin.Context) error {

	queryString := "update san_pham set "
	i := 0
	id, _ := strconv.Atoi(c.PostForm("id"))
	idDanhMuc := c.PostForm("idDM")
	if idDanhMuc != model.EmptyString {
		idDanhMucQr, _ := strconv.Atoi(idDanhMuc)
		i++
		if i > 1 {
			queryString = queryString + ", "
			i--
		}
		queryString = queryString + fmt.Sprintf("id_loaisanpham='%d'", idDanhMucQr)
	}
	tenSP := c.PostForm("tensp")
	if tenSP != model.EmptyString {
		i++
		if i > 1 {
			queryString = queryString + ", "
			i--
		}
		queryString = queryString + fmt.Sprintf("ten_sanpham='%s'", tenSP)
	}
	giaBan := c.PostForm("giaban")
	if giaBan != model.EmptyString {
		i++
		if i > 1 {
			queryString = queryString + ", "
			i--
		}
		giaBanQr, _ := strconv.ParseFloat(giaBan, 64)
		queryString = queryString + fmt.Sprintf("gia_ban='%.2f'", giaBanQr)
	}
	giaNhap := c.PostForm("gianhap")
	if giaNhap != model.EmptyString {
		i++
		if i > 1 {
			queryString = queryString + ", "
			i--
		}
		giaNhapQr, _ := strconv.ParseFloat(giaNhap, 64)
		queryString = queryString + fmt.Sprintf("gia_nhap='%.2f'", giaNhapQr)
	}
	soluong := c.PostForm("soluong")
	if soluong != model.EmptyString {
		i++
		if i > 1 {
			queryString = queryString + ", "
			i--
		}
		soluongQr, _ := strconv.Atoi(soluong)
		queryString = queryString + fmt.Sprintf("so_luong='%d'", soluongQr)
	}
	mota := c.PostForm("mota")
	if mota != model.EmptyString {
		i++
		if i > 1 {
			queryString = queryString + ", "
			i--
		}
		queryString = queryString + fmt.Sprintf("mo_ta='%s'", mota)
	}
	queryString = queryString + fmt.Sprintf(" where id='%d'", id)
	return p.proService.AlterProduct(queryString)
}

func (p ProductService) DeleteSoftProduct(c *gin.Context) error {
	id, _ := strconv.Atoi(c.PostForm("id"))
	status := 0
	return p.proService.DeleteSoftProduct(id, status)
}
