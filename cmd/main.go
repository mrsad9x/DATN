package main

import (
	"DATN/api"
	"DATN/configs"
	"DATN/repository/s3"
	"log"
)

func main() {
	cfg, err := initConfig()
	if err != nil {
		log.Print(err.Error())
		return
	}
	s3Store, err := s3.NewS3Repo(cfg)
	server, err := api.New(cfg, &s3Store)
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = server.Start()

}

func initConfig() (*configs.Server, error) {
	config, err := configs.Init("definition", "common.json")
	if err != nil {
		return nil, err
	}
	return config, nil
}
