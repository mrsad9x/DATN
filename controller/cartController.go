package controller

import "DATN/service"

type CartController struct {
	cartController service.CartService
}

func NewCartController(serv service.CartService) ICartController {
	return CartController{cartController: serv}
}
