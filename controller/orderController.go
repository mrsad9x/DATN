package controller

import "DATN/service"

type OrderController struct {
	orderController service.IOrderService
}

func NewOrderController(order service.OrderService) IOrderController {
	return OrderController{orderController: order}
}
