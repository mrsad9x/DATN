package repository

import (
	"DATN/model"
	"database/sql"
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
	products, err := parseData(data)
	if err != nil {
		return nil, err
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

func (d dbSanPham) GetListProductWithCategories(idDanhMuc int) ([]model.SanPham, error) {
	queryString := fmt.Sprintf("Select * from san_pham where id_DanhMucSP = '%d'", idDanhMuc)
	data, err := d.client.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer data.Close()
	products, err := parseData(data)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (d dbSanPham) SearchProduct(name string) ([]model.SanPham, error) {
	queryString := fmt.Sprintf("Select * from san_pham where tenSP like '%s%s%s'", "%", name, "%")
	data, err := d.client.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer data.Close()

	products, err := parseData(data)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (d dbSanPham) CreateNewProduct(idDM int, tenSP string, giaBan, giaNhap float64, soLuong int, mota string, status int) error {
	queryString := fmt.Sprintf("Insert into san_pham value ('0',%d,'%s',%.2f,%.2f,%d,'%s',%d)", idDM, tenSP, giaBan, giaNhap, soLuong, mota, status)
	return d.client.Exec(queryString)
}

func (d dbSanPham) AlterProduct(id, idDM int, tenSP string, giaBan, giaNhap float64, soLuong int, mota string) error {
	queryString := fmt.Sprintf("Update san_pham set id_danhMucSP = %d, tenSP = '%s', giaBan = %.2f, giaNhap = %.2f, soLuong = %d, moTaSP = '%s' where id = %d", idDM, tenSP, giaBan, giaNhap, soLuong, mota, id)
	return d.client.Exec(queryString)
}

func (d dbSanPham) DeleteSoftProduct(id, status int) error {
	queryString := fmt.Sprintf("Update san_pham set status = %d where id = %d", status, id)
	return d.client.Exec(queryString)
}

func parseData(data *sql.Rows) ([]model.SanPham, error) {
	var products []model.SanPham
	var product model.SanPham
	var id, idDanhmuc, soLuong, status int
	var giaBan, giaNhap float64
	var tensp, mota string

	for data.Next() {
		err := data.Scan(&id, &idDanhmuc, &tensp, &giaBan, &giaNhap, &soLuong, &mota, &status)
		if err != nil {
			return nil, err
		}
		product.Id = id
		product.IdDanhMucSP = idDanhmuc
		product.TenSP = tensp
		product.GiaBan = giaBan
		product.GiaNhap = giaNhap
		product.SoLuong = soLuong
		product.MoTaSP = mota
		product.Status = status

		products = append(products, product)
	}
	return products, nil
}
