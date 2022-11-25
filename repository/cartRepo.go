package repository

type dbCart struct {
	client IDatabase
}

func NewSQLCart(db IDatabase) ICartDB {
	return dbCart{client: db}
}
