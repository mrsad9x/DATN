package model

type NguoiDung struct {
	Id            int    `db:"id"`
	TenNguoiDung  string `db:"ho_ten"`
	TaiKhoan      string `db:"tai_khoan"`
	MatKhau       string `db:"mat_khau"`
	Email         string `db:"email"`
	SoDT          string `db:"so_dien_thoai"`
	DiaChi        string `db:"dia_chi"`
	TrangThai     int    `db:"trang_thai"`
	Role          int    `db:"role"`
	ChiSoTiemNang int    `db:"chi_so_tiem_nang"`
}
