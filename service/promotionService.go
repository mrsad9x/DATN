package service

import "DATN/repository"

type PromotionService struct {
	promoService repository.IPromotionDB
}

func NewPromotionService(repo repository.IPromotionDB) IPromotionService {
	return PromotionService{promoService: repo}
}
