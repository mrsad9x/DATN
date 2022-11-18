package api

import (
	"DATN/configs"
	"DATN/controller"
	"DATN/repository"
	"DATN/service"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	cfg   *configs.Server
	route *gin.Engine
}

func New(cfg *configs.Server) *Server {
	return &Server{
		cfg:   cfg,
		route: gin.Default(),
	}
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
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	userController.SetRouterUserController(s.route)

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
