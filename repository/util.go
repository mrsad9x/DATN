package repository

import (
	"github.com/jmoiron/sqlx"
)

func GetIntFromDataQuery(data *sqlx.Rows) int {
	var index int
	for data.Next() {
		data.Scan(&index)
	}
	return index
}
