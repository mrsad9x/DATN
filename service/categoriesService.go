package service

import "DATN/repository"

type CategoriesService struct {
	cateService repository.ICategoriesDB
}

func NewCateService(repo repository.ICategoriesDB) ICategoriesService {
	return CategoriesService{cateService: repo}
}
