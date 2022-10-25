package service

import "DATN/configs"

type home struct {
	cfg configs.Server
}

func Home() error {
	return nil
}
