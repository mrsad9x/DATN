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
	Login(taiKhoan string) (string, int, error)
	CheckExist(taiKhoan, email string) (bool, int)
	ShowListUser() ([]model.User, error)
	AlterUser(queryString string) error
}

type IProductDB interface {
	GetAllProduct() ([]model.SanPham, error)
	GetOneProduct(id int) (model.SanPham, error)
	GetListProductWithCategories(id int) ([]model.SanPham, error)
	SearchProduct(name string) ([]model.SanPham, error)
	CreateNewProduct(idDM int, tenSP string, giaBan, giaNhap float64, soLuong int, mota string, status int, imgName string) error
	AlterProduct(queryString string) error
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
