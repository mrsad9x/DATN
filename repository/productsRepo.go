package repository

import (
	"DATN/model"
	"database/sql"
	"fmt"
	"strings"
)

type dbProduct struct {
	client IDatabase
}

func NewSQLProduct(db IDatabase) IProductDB {
	return &dbProduct{client: db}
}

func (d dbProduct) GetAllProduct() ([]model.SanPham, error) {
	queryString := "select sp.id, id_loaisanpham, ten_sanpham, gia_ban, gia_nhap, so_luong, mo_ta,trang_thai, a.ten_anh from san_pham sp inner join anh a on sp.id = a.id_sp"
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
	queryString := fmt.Sprintf("select sp.id,ten_sanpham,id_loaisanpham, gia_ban, gia_nhap, so_luong, mo_ta,trang_thai, a.ten_anh from san_pham sp inner join anh a on sp.id = a.id_sp where sp.id='%d'", id)
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
	queryString := fmt.Sprintf("select sp.id,ten_sanpham,id_loaisanpham, gia_ban, gia_nhap, so_luong, mo_ta,trang_thai, a.ten_anh from san_pham sp inner join anh a on sp.id = a.id_sp where id_loaisanpham = '%d'", idDanhMuc)
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
	queryString := fmt.Sprintf("select sp.id,ten_sanpham,id_loaisanpham, gia_ban, gia_nhap, so_luong, mo_ta,trang_thai, a.ten_anh from san_pham sp inner join anh a on sp.id = a.id_sp where ten_sanpham like '%s%s%s'", "%", name, "%")
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

func (d dbProduct) CreateNewProduct(idDM int, tenSP string, giaBan, giaNhap float64, soLuong int, mota string, status int, imgName string) error {
	// check imgName existed
	queryStringSearchImgName := fmt.Sprintf("Select count(*) from anh where ten_anh='%s'", imgName)
	data, err := d.client.QueryOneRow(queryStringSearchImgName)
	if err != nil {
		return err
	}
	quantity := GetIntFromDataQuery(data)
	if quantity != 0 {
		imgName = strings.Replace(imgName, ".", "(1).", 1)
	}
	//queryStringInsertToImageDB := fmt.Sprintf("Insert into anh value('0','%s')", imgName)
	//err = d.client.Exec(queryStringInsertToImageDB)
	//if err != nil {
	//	return err
	//}
	//queryStringGetId := fmt.Sprintf("select id from anh where ten_anh='%s'", imgName)
	//dataId, err := d.client.QueryOneRow(queryStringGetId)
	//if err != nil {
	//	return err
	//}
	//idImg := GetIntFromDataQuery(dataId)
	// insert data to table product
	queryStringInsertToProductDB := fmt.Sprintf("Insert into san_pham value ('0',%d,'%s',%.2f,%.2f,%d,'%s',%d)", idDM, tenSP, giaBan, giaNhap, soLuong, mota, status)
	err = d.client.Exec(queryStringInsertToProductDB)
	if err != nil {
		return err
	}
	// get index newest of product
	queryStringGetNewestProd := fmt.Sprintf("Select count(*) from san_pham")
	dataPro, err := d.client.QueryOneRow(queryStringGetNewestProd)
	if err != nil {
		return err
	}
	idProdNewest := GetIntFromDataQuery(dataPro)

	// insert data to table image
	queryStringInsertToImageDB := fmt.Sprintf("Insert into anh value('0','%d','%s')", idProdNewest, imgName)
	err = d.client.Exec(queryStringInsertToImageDB)
	if err != nil {
		return err
	}
	return nil
}

func (d dbProduct) AlterProduct(queryString string) error {
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
		product.IdCategories = idDanhmuc
		product.ProdName = tensp
		product.Price = giaBan
		product.ImportPrice = giaNhap
		product.Quantity = soLuong
		product.Description = mota
		product.Status = status
		product.PathImg = "static/image/" + pathImg

		products = append(products, product)
	}
	return products, nil
}
