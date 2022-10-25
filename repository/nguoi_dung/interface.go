package nguoi_dung

import "database/sql"

type INguoiDung interface {
	Login(taiKhoan string) (*sql.Rows, error)
	Register(ten, taiKhoan, matKhau, sdt string, status, role, chiSoTN int) error
}
