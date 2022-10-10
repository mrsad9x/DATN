package infrastructure

import (
	"DATN/configs"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type mysql struct {
	config *configs.Database
	client *sql.DB
}

//func NewDB(cfg configs.Database) IDatabase {
//
//	return &mysql{
//		config: &cfg,
//		client: new(sql.DB),
//	}
//}
func NewDBHandle(cfg configs.Database, host string) (IDatabase, error) {
	myclient := &mysql{
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
func (c *mysql) init(host string) (*sql.DB, error) {
	connectInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.config.UserName, c.config.Password, host, c.config.Port, c.config.DBName)
	client, err := sql.Open(c.config.Driver, connectInfo)
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

func (c *mysql) Exec() error {
	queryString := "Select * From demo where QAS is not null "
	data, err := c.client.Query(queryString)
	if err != nil {
		log.Println(err)
		return err
	}
	//for data.Next() {
	//	log.Println(data.Scan())
	//}
	fmt.Println(data)
	return nil
}
