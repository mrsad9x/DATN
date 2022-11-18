package repository

import (
	"DATN/configs"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type Mysql struct {
	config *configs.Database
	client *sqlx.DB
}

func NewDBHandle(cfg configs.Database, host string) (IDatabase, error) {
	myclient := &Mysql{
		config: &cfg,
	}
	svdb, err := myclient.init(host)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	myclient.client = svdb
	return myclient, nil
}
func (c *Mysql) init(host string) (*sqlx.DB, error) {
	connectInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.config.UserName, c.config.Password, host, c.config.Port, c.config.DBName)
	client, err := sqlx.Open(c.config.Driver, connectInfo)
	if err != nil {
		log.Println("Error: ", err.Error())
		return nil, err
	}
	pingErr := client.Ping()
	if pingErr != nil {
		log.Println("Error: ", pingErr.Error())
		return nil, err
	}
	return client, nil
}

func (c *Mysql) Exec(queryString string) error {

	_, err := c.client.Exec(queryString)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (c *Mysql) QueryOneRow(queryString string) (*sqlx.Rows, error) {
	data, err := c.client.Queryx(queryString)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Mysql) Query(queryString string) (*sql.Rows, error) {
	data, err := c.client.Query(queryString)
	if err != nil {
		return nil, err
	}
	return data, nil
}
