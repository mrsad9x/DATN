package service

import "DATN/repository"

type CartService struct {
	cartService repository.ICartDB
}

func NewCartService(repo repository.ICartDB) ICartService {
	return CartService{cartService: repo}
}
