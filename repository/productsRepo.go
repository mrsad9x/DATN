package repository

import (
	"DATN/model"
	"database/sql"
	"fmt"
)

type dbProduct struct {
	client IDatabase
}

func NewSQLProduct(db IDatabase) IProductDB {
	return &dbProduct{client: db}
}

func (d dbProduct) GetAllProduct() ([]model.SanPham, error) {
	queryString := "select sp.id, id_loaisanpham, ten_sanpham, gia_ban, gia_nhap, so_luong, mo_ta,trang_thai, a.ten_anh from san_pham sp inner join anh a on sp.id_anh = a.id"
	data, err := d.client.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer data.Close()
	products, err := ParseDataProduct(data)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (d dbProduct) GetOneProduct(id int) (model.SanPham, error) {
	queryString := fmt.Sprintf("select sp.id,ten_sanpham,id_loaisanpham, gia_ban, gia_nhap, so_luong, mo_ta,trang_thai, a.ten_anh from san_pham sp inner join anh a on sp.id_anh = a.id where sp.id='%d'", id)
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

func (d dbProduct) GetListProductWithCategories(idDanhMuc int) ([]model.SanPham, error) {
	queryString := fmt.Sprintf("select sp.id,ten_sanpham,id_loaisanpham, gia_ban, gia_nhap, so_luong, mo_ta,trang_thai, a.ten_anh from san_pham sp inner join anh a on sp.id_anh = a.id where id_loaisanpham = '%d'", idDanhMuc)
	data, err := d.client.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer data.Close()
	products, err := ParseDataProduct(data)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (d dbProduct) SearchProduct(name string) ([]model.SanPham, error) {
	queryString := fmt.Sprintf("select sp.id,ten_sanpham,id_loaisanpham, gia_ban, gia_nhap, so_luong, mo_ta,trang_thai, a.ten_anh from san_pham sp inner join anh a on sp.id_anh = a.id where tenSP like '%s%s%s'", "%", name, "%")
	data, err := d.client.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer data.Close()

	products, err := ParseDataProduct(data)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (d dbProduct) CreateNewProduct(idDM int, tenSP string, giaBan, giaNhap float64, soLuong int, mota string, status int) error {
	queryString := fmt.Sprintf("Insert into san_pham value ('0',%d,'%s',%.2f,%.2f,%d,'%s',%d)", idDM, tenSP, giaBan, giaNhap, soLuong, mota, status)
	return d.client.Exec(queryString)
}

func (d dbProduct) AlterProduct(id, idDM int, tenSP string, giaBan, giaNhap float64, soLuong int, mota string) error {
	queryString := fmt.Sprintf("Update san_pham set id_danhMucSP = %d, tenSP = '%s', giaBan = %.2f, giaNhap = %.2f, soLuong = %d, moTaSP = '%s' where id = %d", idDM, tenSP, giaBan, giaNhap, soLuong, mota, id)
	return d.client.Exec(queryString)
}

func (d dbProduct) DeleteSoftProduct(id, status int) error {
	queryString := fmt.Sprintf("Update san_pham set status = %d where id = %d", status, id)
	return d.client.Exec(queryString)
}

func ParseDataProduct(data *sql.Rows) ([]model.SanPham, error) {
	var products []model.SanPham
	var product model.SanPham
	var id, idDanhmuc, soLuong, status int
	var giaBan, giaNhap float64
	var tensp, mota, pathImg string

	for data.Next() {
		err := data.Scan(&id, &tensp, &idDanhmuc, &giaBan, &giaNhap, &soLuong, &mota, &status, &pathImg)
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
		product.PathImg = "static/image/" + pathImg

		products = append(products, product)
	}
	return products, nil
}
