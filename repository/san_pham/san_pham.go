package san_pham

import "DATN/repository"

type dbSanPham struct {
	client repository.IDatabase
}

func NewSQLSanPham(db repository.IDatabase) ISanPham {
	return &dbSanPham{client: db}
}
