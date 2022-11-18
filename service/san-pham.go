package service

import "DATN/repository"

type ProductService struct {
	ProService repository.ISanPham
}

func NewProducService(repo repository.ISanPham) IProductService {
	return ProductService{ProService: repo}
}

func (p ProductService) LaySanPham() {

}

func (p ProductService) LaySanPhamCoDieuKien() {

}
