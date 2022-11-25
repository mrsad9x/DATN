package repository

type dbCategories struct {
	client IDatabase
}

func NewSQLCategories(db IDatabase) ICategoriesDB {
	return dbCategories{client: db}
}
