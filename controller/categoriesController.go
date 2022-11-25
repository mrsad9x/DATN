package controller

import "DATN/service"

type CategoriesController struct {
	cateController service.CategoriesService
}

func NewCategoriesController(serv service.CategoriesService) ICategoriesController {
	return CategoriesController{cateController: serv}
}
