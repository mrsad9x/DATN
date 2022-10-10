package configs

type Database struct {
	Port     int    `mapstructure:"DB_PORT"`
	DBName   string `mapstructure:"DB_NAME"`
	UserName string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Driver   string `mapstructure:"DB_DRIVER_NAME"`
}
