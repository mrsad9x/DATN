package api

import (
	"DATN/configs"
	"DATN/repository"
	nguoiDung "DATN/repository/nguoi_dung"
	"DATN/service"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Server struct {
	cfg *configs.Server
}

func New(cfg *configs.Server) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Start() error {

	db, err := repository.NewDBHandle(s.cfg.Database, "localhost")
	if err != nil {
		log.Println(err)
		return err
	}
	userRepo := nguoiDung.NewSQLNguoiDung(db)
	//dbSanPham := sanPham.NewSQLSanPham(db)
	//err = userRepo.Login("admin", "123456123")
	//if err != nil {
	//	return err
	//}
	userService := service.NewUserService(userRepo)
	//_ = service.NewUserService(userRepo)

	err = userService.Register("Duc Xuan", "mrsad9x", "123456", "0375686987", 1, 1, 1)
	if err != nil {
		return err
	}
	time.Sleep(5 * time.Second)
	err = userService.Login("mrsad9x", "123456")
	if err != nil {
		fmt.Println("login fail")
	}
	//err = http.ListenAndServe(model.PortConn, nil)
	//if err != nil {
	//	log.Println(err)
	//}
	//http.HandleFunc("/", HandleRequest)

	return nil
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {

	switch r.RequestURI {
	case "/":

	case "/home":

	}

}
