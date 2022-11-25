package service

import "DATN/repository"

type OrderService struct {
	orderService repository.IOrderDB
}

func NewOrderService(repo repository.IOrderDB) IOrderService {
	return OrderService{orderService: repo}
}
