package service

import (
	"DATN/repository"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type PromotionService struct {
	promoService repository.IPromotionDB
}

func NewPromotionService(repo repository.IPromotionDB) IPromotionService {
	return PromotionService{promoService: repo}
}

func (p PromotionService) CreatePromotion(c *gin.Context) error {
	maGiamGia := c.PostForm("magiamgia")
	timeCreate := time.Now().Format("20060102150405")
	dayApply, _ := strconv.Atoi(c.PostForm("timeapdung"))
	pencentApply, _ := strconv.Atoi(c.PostForm("phantram"))
	mota := c.PostForm("mota")
	rankApply, _ := strconv.Atoi(c.PostForm("rankapdung"))

	return p.promoService.CreatePromotion(maGiamGia, timeCreate, mota, dayApply, pencentApply, rankApply)
}

func (p PromotionService) AlterPromotion(c *gin.Context) error {
	return nil
}

func (p PromotionService) DeletePromotion(c *gin.Context) error {
	return nil
}
