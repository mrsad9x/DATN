package repository

import (
	"DATN/model"
	"fmt"
	"log"
)

type dbNguoiDung struct {
	client IDatabase
}

func NewSQLNguoiDung(db IDatabase) INguoiDungDB {
	return &dbNguoiDung{client: db}
}

func (d dbNguoiDung) Login(taiKhoan string) (string, error) {
	queryCommand := fmt.Sprintf("Select * from nguoi_dung where taiKhoan = '%s'", taiKhoan)
	data, err := d.client.QueryOneRow(queryCommand)
	if err != nil {
		return "", err
	}
	var user model.NguoiDung
	for data.Next() {
		err = data.StructScan(&user)
		if err != nil {
			return "", err
		}
	}
	fmt.Println(user.TaiKhoan)
	return user.MatKhau, nil
}

func (d dbNguoiDung) Register(ten, taiKhoan, matKhau, sdt, email, diaChi string, status, role, chiSoTN int) error {
	queryCommand := fmt.Sprintf("Insert into nguoi_dung value('0','%s','%s','%s','%s','%s','%s','%d','%d','%d')", ten, taiKhoan, matKhau, sdt, email, diaChi, status, role, chiSoTN)
	err := d.client.Exec(queryCommand)
	if err != nil {
		return err
	}
	queryCommand = fmt.Sprintf("Select * from nguoi_dung where taiKhoan = '%s' or email = '%s'", taiKhoan, email)
	data, err := d.client.QueryOneRow(queryCommand)
	if err != nil {
		log.Println(err.Error())
	}
	var nd model.NguoiDung
	defer data.Close()
	data.Next()
	{
		err = data.StructScan(&nd)
		if err != nil {
			log.Println(err.Error())
		}
	}
	queryCommand = fmt.Sprintf("Insert into role_user value('0','%d','%d')", nd.Id, nd.Role)
	err = d.client.Exec(queryCommand)
	if err != nil {
		return err
	}
	return nil
}

func (d dbNguoiDung) CheckExist(taiKhoan, email string) (bool, int) {
	var nd model.NguoiDung
	queryCommand := fmt.Sprintf("Select * from nguoi_dung where taiKhoan = '%s' or email = '%s'", taiKhoan, email)
	data, err := d.client.QueryOneRow(queryCommand)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer data.Close()
	data.Next()
	{
		err = data.StructScan(&nd)
		if nd.TaiKhoan == taiKhoan {
			return true, 1
		}
		if nd.Email == email {
			return true, 2
		}
	}
	return false, 0
}
