package repository

import "fmt"

type dbPromotion struct {
	client IDatabase
}

func NewSQLPromotion(db IDatabase) IPromotionDB {
	return dbPromotion{client: db}
}

func (db dbPromotion) CreatePromotion(maGG, timeCreate, mota string, dayApply, percentApp, rankApp int) error {
	queryString := fmt.Sprintf("insert into ma_giam_gia value ('0','%s','%s',%d,%d,'%s',%d)", maGG, timeCreate, dayApply, percentApp, mota, rankApp)
	return db.client.Exec(queryString)
}
