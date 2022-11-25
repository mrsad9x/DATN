package model

type SanPham struct {
	Id          int     `db:"id"`
	IdDanhMucSP int     `db:"id_loaisanpham"`
	TenSP       string  `db:"ten_sanpham"`
	GiaBan      float64 `db:"gia_ban"`
	GiaNhap     float64 `db:"gia_nhap"`
	SoLuong     int     `db:"so_luong"`
	MoTaSP      string  `db:"mo_ta"`
	Status      int     `db:"trang_thai"`
	PathImg     string  `db:"ten_anh"`
}
