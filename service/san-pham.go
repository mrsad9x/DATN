package service

import (
	"DATN/model"
	"DATN/repository"
)

type ProductService struct {
	ProService repository.ISanPham
}

func NewProducService(repo repository.ISanPham) IProductService {
	return ProductService{ProService: repo}
}

func (p ProductService) GetAllProduct() ([]model.SanPham, error) {
	return p.ProService.GetAllProduct()
}

func (p ProductService) GetOneProduct(id int) (model.SanPham, error) {
	return p.ProService.GetOneProduct(id)
}

func (p ProductService) GetListProductWithCategories(id int) ([]model.SanPham, error) {
	return p.ProService.GetListProductWithCategories(id)
}

func (p ProductService) SearchProduct(name string) ([]model.SanPham, error) {
	return p.ProService.SearchProduct(name)
}

func (p ProductService) CreateNewProduct(idDM int, tenSP string, giaBan, giaNhap float64, soLuong int, mota string, status int) error {
	return p.ProService.CreateNewProduct(idDM, tenSP, giaBan, giaNhap, soLuong, mota, status)
}

func (p ProductService) AlterProduct(id, idDM int, tenSP string, giaBan, giaNhap float64, soLuong int, mota string) error {
	return p.ProService.AlterProduct(id, idDM, tenSP, giaBan, giaNhap, soLuong, mota)
}

func (p ProductService) DeleteSoftProduct(id int) error {
	status := 0
	return p.ProService.DeleteSoftProduct(id, status)
}
