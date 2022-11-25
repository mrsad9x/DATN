package controller

import "DATN/service"

type PromotionController struct {
	promoController service.IPromotionService
}

func NewPromotionController(serv service.IPromotionService) IPromotionController {
	return PromotionController{promoController: serv}
}
