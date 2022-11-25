package repository

type dbPromotion struct {
	client IDatabase
}

func NewSQLPromotion(db IDatabase) IPromotionDB {
	return dbPromotion{client: db}
}
