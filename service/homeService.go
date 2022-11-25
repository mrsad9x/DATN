package service

import (
	"DATN/model"
	"DATN/repository"
)

type HomeService struct {
	homeService repository.IHomeDB
}

func NewHomeService(repo repository.IHomeDB) IHomeService {
	return HomeService{homeService: repo}
}

func (h HomeService) Home() ([]model.SanPham, error) {
	return h.homeService.Home()
}
