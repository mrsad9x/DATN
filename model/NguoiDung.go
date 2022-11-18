package model

type NguoiDung struct {
	Id            int    `db:"id"`
	TenNguoiDung  string `db:"tenND"`
	TaiKhoan      string `db:"taiKhoan"`
	MatKhau       string `db:"matKhau"`
	Email         string `db:"email"`
	SoDT          string `db:"soDT"`
	DiaChi        string `db:"diaChi"`
	TrangThai     int    `db:"trangThai"`
	Role          int    `db:"role"`
	ChiSoTiemNang int    `db:"chiSoTiemNang"`
}
