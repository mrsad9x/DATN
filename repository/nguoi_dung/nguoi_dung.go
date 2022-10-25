package nguoi_dung

import (
	"DATN/repository"
	"database/sql"
	"fmt"
)

type dbNguoiDung struct {
	client repository.IDatabase
}

func NewSQLNguoiDung(db repository.IDatabase) INguoiDung {
	return &dbNguoiDung{client: db}
}

func (d dbNguoiDung) Login(taiKhoan string) (*sql.Rows, error) {
	queryCommand := fmt.Sprintf("Select * from nguoi_dung  where taiKhoan = '%s'", taiKhoan)
	data, err := d.client.QueryRow(queryCommand)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (d dbNguoiDung) Register(ten, taiKhoan, matKhau string, sdt string, status, role, chiSoTN int) error {
	queryCommand := fmt.Sprintf("Insert into nguoi_dung value('0','%s','%s','%s','%s','%d','%d','%d')", ten, taiKhoan, matKhau, sdt, status, role, chiSoTN)
	err := d.client.Exec(queryCommand)
	if err != nil {
		return err
	}
	return nil
}
