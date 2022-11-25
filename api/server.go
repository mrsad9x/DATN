package api

import (
	"DATN/configs"
	"DATN/controller"
	"DATN/repository"
	"DATN/service"
	"DATN/token"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	cfg        *configs.Server
	route      *gin.Engine
	tokenMaker token.Maker
}

func New(cfg *configs.Server) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(cfg.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("can not create token maker: %w", err)
	}
	return &Server{
		cfg:        cfg,
		route:      gin.Default(),
		tokenMaker: tokenMaker,
	}, nil
}

func (s *Server) Start() error {
	port := "8080"
	db, err := repository.NewDBHandle(s.cfg.Database, "localhost")
	if err != nil {
		log.Println(err)
		return err
	}

	err = s.route.SetTrustedProxies(nil)

	userRepo := repository.NewSQLNguoiDung(db)
	userService := service.NewUserService(userRepo, s.tokenMaker, s.cfg)
	userController := controller.NewUserController(userService)
	userController.SetRouterUserController(s.route)

	prodRepo := repository.NewSQLProduct(db)
	prodService := service.NewProducService(prodRepo)
	prodController := controller.NewProductController(prodService)
	prodController.SetRouterSanPhamController(s.route)

	homeRepo := repository.NewSQLHome(db)
	homeService := service.NewHomeService(homeRepo)
	homeController := controller.NewHomeController(homeService)
	homeController.SetRouterHomeController(s.route)

	s.route.LoadHTMLGlob("templates/*.html")
	s.route.Use(cors.Default())
	err = s.route.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		return err
	}

	//err = userService.Register("Duc Xuan", "mrsad9x", "123456", "0375686987", "mr.sad.9x@gmail.com", "", 1, 1, 1)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return err
	//}

	return nil
}
