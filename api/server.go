package api

import (
	"DATN/configs"
	"DATN/controller"
	"DATN/repository"
	"DATN/repository/s3"
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
	s3Store    *s3.IS3Repo
}

func New(cfg *configs.Server, s3Store *s3.IS3Repo) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(cfg.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("can not create token maker: %w", err)
	}
	return &Server{
		cfg:        cfg,
		route:      gin.Default(),
		tokenMaker: tokenMaker,
		s3Store:    s3Store,
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

	userRepo := repository.NewSQLUser(db)
	userService := service.NewUserService(userRepo, s.tokenMaker, s.cfg)
	userController := controller.NewUserController(userService)
	userController.SetRouterUserController(s.route)

	prodRepo := repository.NewSQLProduct(db)
	prodService := service.NewProducService(prodRepo, s.s3Store)
	prodController := controller.NewProductController(prodService, userController)
	prodController.SetRouterSanPhamController(s.route)

	homeRepo := repository.NewSQLHome(db)
	homeService := service.NewHomeService(homeRepo)
	homeController := controller.NewHomeController(homeService)
	homeController.SetRouterHomeController(s.route)

	promotionRepo := repository.NewSQLPromotion(db)
	promotionService := service.NewPromotionService(promotionRepo)
	promotionController := controller.NewPromotionController(promotionService, userController)
	promotionController.SetRouterPromotionController(s.route)

	s.route.Static("/static", "./templates/static")
	s.route.LoadHTMLGlob("templates/*.*")

	s.route.Use(cors.Default())
	err = s.route.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		return err
	}

	return nil
}
