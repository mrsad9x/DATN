package service

import (
	"DATN/repository"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo repository.INguoiDungDB
}

func NewUserService(repo repository.INguoiDungDB) IUserService {
	return UserService{UserRepo: repo}
}
func (c UserService) Register(ten, taiKhoan, matKhau, sdt, email, diaChi string, status, role, chiSoTN int) error {
	exist, check := c.UserRepo.CheckExist(taiKhoan, email)
	if exist && check == 1 {
		return fmt.Errorf("user name existed")
	}
	if exist && check == 2 {
		return fmt.Errorf("email existed")
	}
	pass, err := hashPassword(matKhau)
	if err != nil {
		return err
	}
	err = c.UserRepo.Register(ten, taiKhoan, pass, sdt, email, diaChi, status, role, chiSoTN)
	if err != nil {
		return err
	}
	return nil
}

func (c UserService) Login(taiKhoan, matKhau string) error {
	passHash, err := c.UserRepo.Login(taiKhoan)
	if err != nil {
		return err
	}
	result := checkPassword(matKhau, passHash)
	if result {
		fmt.Println("true")
	}
	return nil
}

func hashPassword(matKhau string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(matKhau), 14)
	return string(bytes), err
}

func checkPassword(passWord, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(passWord))
	return err == nil
}
