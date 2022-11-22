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

type INguoiDungDB interface {
	Register(ten, taiKhoan, matKhau, sdt, email, diaChi string, status, role, chiSoTN int) error
	Login(taiKhoan string) (string, error)
	CheckExist(taiKhoan, email string) (bool, int)
}

type ISanPham interface {
	GetAllProduct() ([]model.SanPham, error)
	GetOneProduct(id int) (model.SanPham, error)
}
