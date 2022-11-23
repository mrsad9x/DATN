package service

import "DATN/model"

type IUserService interface {
	Register(ten, taiKhoan, matKhau, sdt, email, diaChi string, status, role, chiSoTN int) error
	Login(taiKhoan, matKhau string) error
}

type IProductService interface {
	GetAllProduct() ([]model.SanPham, error)
	GetOneProduct(id int) (model.SanPham, error)
	GetListProductWithCategories(id int) ([]model.SanPham, error)
	SearchProduct(name string) ([]model.SanPham, error)
	CreateNewProduct(idDM int, tenSP string, giaBan, giaNhap float64, soLuong int, mota string, status int) error
	AlterProduct(id, idDM int, tenSP string, giaBan, giaNhap float64, soLuong int, mota string) error
	DeleteSoftProduct(id int) error
}
