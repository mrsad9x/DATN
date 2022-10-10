package model

type NguoiDung struct {
	id            int    `db:"id"`
	tenNguoiDung  string `db:"ten-nd"`
	taiKhoan      string
	matKhau       string
	soDT          int
	trangThai     rune
	role          rune
	chiSoTiemNang rune
}
