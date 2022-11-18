package repository

import "fmt"

type dbSanPham struct {
	client IDatabase
}

func NewSQLSanPham(db IDatabase) ISanPham {
	return &dbSanPham{client: db}
}

func (d dbSanPham) GetAllProduct() {
	queryString := "Select * form san_pham"
	data, err := d.client.Query(queryString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
