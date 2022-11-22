package model

type SanPham struct {
	Id          int     `db:"id"`
	IdDanhMucSP int     `db:"id_danhMucSP"`
	TenSP       string  `db:"tenSP"`
	GiaBan      float64 `db:"giaBan"`
	GiaNhap     float64 `db:"giaNhap"`
	SoLuong     int     `db:"soLuong"`
	MoTaSP      string  `db:"moTaSP"`
}
