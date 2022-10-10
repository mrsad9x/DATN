package main

import (
	"DATN/api"
	"DATN/configs"
	"log"
)

func main() {
	cfg, err := initConfig()
	if err != nil {
		log.Print(err.Error())
		return
	}
	server := api.New(cfg)
	err = server.Start()

}

func initConfig() (*configs.Server, error) {
	config, err := configs.Init("definition", "common.json")
	if err != nil {
		return nil, err
	}
	return config, nil
}
