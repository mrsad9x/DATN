package service

type IUserService interface {
	Register(ten, taiKhoan, matKhau, sdt, email, diaChi string, status, role, chiSoTN int) error
	Login(taiKhoan, matKhau string) error
}

type IProductService interface {
	LaySanPham()
	LaySanPhamCoDieuKien()
}
