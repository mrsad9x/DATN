package configs

import "time"

type Database struct {
	Port     int    `mapstructure:"DB_PORT"`
	DBName   string `mapstructure:"DB_NAME"`
	UserName string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Driver   string `mapstructure:"DB_DRIVER_NAME"`
}

type Token struct {
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}
