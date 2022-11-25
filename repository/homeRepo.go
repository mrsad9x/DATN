package repository

import (
	"DATN/model"
	"fmt"
)

type dbHome struct {
	cllient IDatabase
}

func NewSQLHome(db IDatabase) IHomeDB {
	return &dbHome{cllient: db}
}

func (d dbHome) Home() ([]model.SanPham, error) {
	queryString := fmt.Sprintf("select sp.id, sp.ten_sanpham,sp.gia_ban, sp.gia_nhap, sp.so_luong, sp.mo_ta, sp.trang_thai from san_pham sp inner join don_hang dh on sp.id = dh.id_sanpham where sp.trang_thai= 1 order by dh.id_sanpham desc limit 8")
	data, err := d.cllient.Query(queryString)
	if err != nil {
		return nil, err
	}
	listProduct, err := ParseDataProduct(data)
	if err != nil {
		return nil, err
	}
	return listProduct, nil
}
