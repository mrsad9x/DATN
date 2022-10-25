package service

import (
	"DATN/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo repository.INguoiDungDB
}

func NewUserService(repo repository.INguoiDungDB) IUser {
	return UserService{UserRepo: repo}
}
func (c UserService) Register(ten, taiKhoan, matKhau, sdt string, status, role, chiSoTN int) error {
	pass, err := hashPassword(matKhau)
	if err != nil {
		return err
	}
	err = c.UserRepo.Register(ten, taiKhoan, pass, sdt, status, role, chiSoTN)
	if err != nil {
		return err
	}
	return nil
}

func (c UserService) Login(taiKhoan, matKhau string) error {
	pass, err := c.UserRepo.Login(taiKhoan)
	if err != nil {
		return err
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
