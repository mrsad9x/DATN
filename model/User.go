package model

import "time"

type User struct {
	Id       int    `db:"id"`
	FullName string `db:"ho_ten"`
	UserName string `db:"tai_khoan"`
	Password string `db:"mat_khau"`
	Email    string `db:"email"`
	Phone    string `db:"so_dien_thoai"`
	Address  string `db:"dia_chi"`
	Status   int    `db:"trang_thai"`
	Role     int    `db:"role"`
	Rank     int    `db:"chi_so_tiem_nang"`
}

type ResponseUser struct {
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}
