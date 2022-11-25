package service

import (
	"DATN/model"
	"DATN/repository"
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

func (p ProductService) GetOneProduct(id int) (model.SanPham, error) {
	return p.proService.GetOneProduct(id)
}

func (p ProductService) GetListProductWithCategories(id int) ([]model.SanPham, error) {
	return p.proService.GetListProductWithCategories(id)
}

func (p ProductService) SearchProduct(name string) ([]model.SanPham, error) {
	return p.proService.SearchProduct(name)
}

func (p ProductService) CreateNewProduct(idDM int, tenSP string, giaBan, giaNhap float64, soLuong int, mota string, status int) error {
	return p.proService.CreateNewProduct(idDM, tenSP, giaBan, giaNhap, soLuong, mota, status)
}

func (p ProductService) AlterProduct(id, idDM int, tenSP string, giaBan, giaNhap float64, soLuong int, mota string) error {
	return p.proService.AlterProduct(id, idDM, tenSP, giaBan, giaNhap, soLuong, mota)
}

func (p ProductService) DeleteSoftProduct(id int) error {
	status := 0
	return p.proService.DeleteSoftProduct(id, status)
}
