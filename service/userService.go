package service

import (
	"DATN/configs"
	"DATN/model"
	"DATN/repository"
	"DATN/token"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo repository.IUserDB
	token    token.Maker
	cfg      *configs.Server
}

func NewUserService(repo repository.IUserDB, token token.Maker, cfg *configs.Server) IUserService {
	return UserService{UserRepo: repo, token: token, cfg: cfg}
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

func (c UserService) Login(taiKhoan, matKhau string) (string, error) {
	passHash, role, err := c.UserRepo.Login(taiKhoan)
	if err != nil {
		return model.EmptyString, err
	}
	result := checkPassword(matKhau, passHash)
	createToken, _, err := c.token.CreateToken(taiKhoan, role, c.cfg.AccessTokenDuration)
	if err != nil {
		return model.EmptyString, err
	}
	if !result {
		return model.EmptyString, fmt.Errorf("login fail")
	}

	return createToken, nil
}

func hashPassword(matKhau string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(matKhau), 14)
	return string(bytes), err
}

func checkPassword(passWord, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(passWord))
	return err == nil
}
