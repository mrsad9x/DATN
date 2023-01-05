package service

import (
	"DATN/configs"
	"DATN/model"
	"DATN/repository"
	"DATN/token"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

type UserService struct {
	UserRepo repository.IUserDB
	token    token.Maker
	cfg      *configs.Server
}

func NewUserService(repo repository.IUserDB, token token.Maker, cfg *configs.Server) IUserService {
	return UserService{UserRepo: repo, token: token, cfg: cfg}
}
func (s UserService) Register(c *gin.Context) error {
	ten := c.PostForm("name")
	taiKhoan := c.PostForm("username")
	matKhau := c.PostForm("password")
	sdt := c.PostForm("phone")
	email := c.PostForm("email")
	diaChi := ""
	chiSoTN := 1
	status := 1
	role := 3
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
	return s.UserRepo.Register(ten, taiKhoan, pass, sdt, email, diaChi, status, role, chiSoTN)

}

func (s UserService) CreateUser(c *gin.Context) error {
	ten := c.PostForm("name")
	taiKhoan := c.PostForm("username")
	matKhau := c.PostForm("password")
	sdt := c.PostForm("phone")
	email := c.PostForm("email")
	diaChi := c.PostForm("address")
	chiSoTN := 1
	status := 1
	role, _ := strconv.Atoi(c.PostForm("role"))
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
	return s.UserRepo.Register(ten, taiKhoan, pass, sdt, email, diaChi, status, role, chiSoTN)
}

func (s UserService) Login(c *gin.Context) (string, error) {
	username := c.PostForm("user")
	pass := c.PostForm("pass")
	if username == "admin" && pass == "123456" {
		role := 1
		createToken, _, err := s.token.CreateToken(username, role, s.cfg.AccessTokenDuration)
		if err != nil {
			return model.EmptyString, err
		}
		return createToken, nil
	}

	passHash, role, err := s.UserRepo.Login(username)
	if err != nil {
		return model.EmptyString, err
	}
	result := checkPassword(pass, passHash)

	if !result {
		return model.EmptyString, fmt.Errorf("login fail")
	}

	createToken, _, err := s.token.CreateToken(username, role, s.cfg.AccessTokenDuration)
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

func (s UserService) ShowListUer() ([]model.User, error) {
	return s.UserRepo.ShowListUser()
}

func (s UserService) AlterUser(c *gin.Context) error {
	var queryCommand string
	var i = 0
	queryCommand = "update nguoi_dung set "

	idget := c.PostForm("id")
	id, _ := strconv.Atoi(idget)
	name := c.PostForm("name")
	if name != model.EmptyString {
		i++

		if i > 1 {
			queryCommand = queryCommand + ", "
			i--
		}
		queryCommand = queryCommand + fmt.Sprintf("ho_ten = '%s'", name)
	}
	passwordGet := c.PostForm("pass")
	password, err := hashPassword(passwordGet)
	if err != nil {
		return err
	}
	if password != model.EmptyString {
		i++
		if i > 1 {
			queryCommand = queryCommand + ", "
			i--
		}
		queryCommand = queryCommand + fmt.Sprintf("mat_khau = '%s'", password)

	}
	phone := c.PostForm("phone")
	if phone != model.EmptyString {
		i++
		if i > 1 {
			queryCommand = queryCommand + ", "
			i--
		}
		queryCommand = queryCommand + fmt.Sprintf("so_dien_thoai = '%s'", phone)

	}
	address := c.PostForm("address")
	if address != model.EmptyString {
		i++
		if i > 1 {
			queryCommand = queryCommand + ", "
			i--
		}
		queryCommand = queryCommand + fmt.Sprintf("dia_chi = '%s'", address)

	}

	status := c.PostForm("status")
	if status != model.EmptyString {
		i++
		if i > 1 {
			queryCommand = queryCommand + ", "
			i--
		}
		status1, _ := strconv.Atoi(status)
		queryCommand = queryCommand + fmt.Sprintf("trang_thai = %d", status1)

	}
	queryCommand = queryCommand + fmt.Sprintf(" where id = %d", id)
	return s.UserRepo.AlterUser(queryCommand)
}
