package service

import "DATN/model"

type IUserService interface {
	Register(ten, taiKhoan, matKhau, sdt, email, diaChi string, status, role, chiSoTN int) error
	Login(taiKhoan, matKhau string) error
}

type IProductService interface {
	GetAllProduct() ([]model.SanPham, error)
	GetOneProduct(id int) (model.SanPham, error)
}
