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
func (s UserService) Register(ten, taiKhoan, matKhau, sdt, email, diaChi string, status, role, chiSoTN int) error {
	exist, check := s.UserRepo.CheckExist(taiKhoan, email)
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
	err = s.UserRepo.Register(ten, taiKhoan, pass, sdt, email, diaChi, status, role, chiSoTN)
	if err != nil {
		return err
	}
	return nil
}

func (s UserService) Login(taiKhoan, matKhau string) (string, error) {
	passHash, role, err := s.UserRepo.Login(taiKhoan)
	if err != nil {
		return model.EmptyString, err
	}
	result := checkPassword(matKhau, passHash)

	if !result {
		return model.EmptyString, fmt.Errorf("login fail")
	}

	createToken, _, err := s.token.CreateToken(taiKhoan, role, s.cfg.AccessTokenDuration)
	if err != nil {
		return model.EmptyString, err
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

func (s UserService) CheckRoles(token string) (int, error) {
	payload, err := s.token.VerifyToken(token)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return payload.RoleUser, nil
}
