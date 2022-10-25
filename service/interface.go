package service

type IUser interface {
	Register(ten, taiKhoan, matKhau, sdt string, status, role, chiSoTN int) error
	Login(taiKhoan, matKhau string) error
}
