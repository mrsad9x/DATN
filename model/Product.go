package model

type SanPham struct {
	Id           int     `db:"id"`
	IdCategories int     `db:"id_loaisanpham"`
	ProdName     string  `db:"ten_sanpham"`
	Price        float64 `db:"gia_ban"`
	ImportPrice  float64 `db:"gia_nhap"`
	Quantity     int     `db:"so_luong"`
	Description  string  `db:"mo_ta"`
	Status       int     `db:"trang_thai"`
	PathImg      string  `db:"ten_anh"`
}
