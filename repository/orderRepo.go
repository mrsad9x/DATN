package repository

type dbOrder struct {
	client IDatabase
}

func NewSQLOrder(db IDatabase) IOrderDB {
	return dbOrder{client: db}
}
