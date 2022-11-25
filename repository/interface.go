package repository

import (
	"DATN/model"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type IDatabase interface {
	Exec(queryString string) error
	QueryOneRow(queryString string) (*sqlx.Rows, error)
	Query(queryString string) (*sql.Rows, error)
}

type IUserDB interface {
	Register(ten, taiKhoan, matKhau, sdt, email, diaChi string, status, role, chiSoTN int) error
	Login(taiKhoan string) (string, error)
	CheckExist(taiKhoan, email string) (bool, int)
}

type IProductDB interface {
	GetAllProduct() ([]model.SanPham, error)
	GetOneProduct(id int) (model.SanPham, error)
	GetListProductWithCategories(id int) ([]model.SanPham, error)
	SearchProduct(name string) ([]model.SanPham, error)
	CreateNewProduct(idDM int, tenSP string, giaBan, giaNhap float64, soLuong int, mota string, status int) error
	AlterProduct(id, idDM int, tenSP string, giaBan, giaNhap float64, soLuong int, mota string) error
	DeleteSoftProduct(id, status int) error
}

type IHomeDB interface {
	Home() ([]model.SanPham, error)
}

type ICategoriesDB interface {
}

type ICartDB interface {
}

type IOrderDB interface {
}

type IPromotionDB interface {
}
