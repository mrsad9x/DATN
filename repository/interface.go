package repository

import "database/sql"

type IDatabase interface {
	Exec(queryString string) error
	QueryRow(queryString string) (*sql.Rows, error)
}

type INguoiDungDB interface {
	Register(ten, taiKhoan, matKhau, sdt string, status, role, chiSoTN int) error
	Login(taiKhoan string) (*sql.Rows, error)
}
