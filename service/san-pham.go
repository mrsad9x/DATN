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
