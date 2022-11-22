package repository

import (
	"DATN/model"
	"fmt"
)

type dbSanPham struct {
	client IDatabase
}

func NewSQLSanPham(db IDatabase) ISanPham {
	return &dbSanPham{client: db}
}

func (d dbSanPham) GetAllProduct() ([]model.SanPham, error) {
	queryString := "Select * from san_pham"
	data, err := d.client.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer data.Close()
	var products []model.SanPham
	var product model.SanPham
	var id, id_danhmuc, soLuong int
	var giaBan, giaNhap float64
	var tensp, mota string

	for data.Next() {
		err = data.Scan(&id, &id_danhmuc, &tensp, &giaBan, &giaNhap, &soLuong, &mota)
		if err != nil {
			return nil, err
		}
		product.Id = id
		product.IdDanhMucSP = id_danhmuc
		product.GiaBan = giaBan
		product.GiaNhap = giaNhap
		product.SoLuong = soLuong
		product.MoTaSP = mota

		products = append(products, product)
	}
	return products, nil
}

func (d dbSanPham) GetOneProduct(id int) (model.SanPham, error) {
	queryString := fmt.Sprintf("Select * from san_pham where id='%d'", id)
	data, err := d.client.QueryOneRow(queryString)
	var product model.SanPham
	defer data.Close()
	data.Next()
	{
		err = data.StructScan(&product)
		if err != nil {
			return product, err
		}
	}
	return product, nil
}
