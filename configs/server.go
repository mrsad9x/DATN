package configs

import (
	"github.com/spf13/viper"
)

type Server struct {
	Database     `mapstructure:"Database"`
	Token        `mapstructure:"Token"`
	AccessKeyAWS `mapstructure:"AccessKeyAWS"`
}

func Init(path, fileName string) (*Server, error) {
	cfg := new(Server)
	viper.AddConfigPath(path)
	viper.SetConfigName(fileName)
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

type AccessKeyAWS struct {
	Id     string `mapstructure:"ID_Key"`
	Secret string `mapstructure:"SECRET_KEY"`
	Region string `mapstructure:"AWS_REGION"`
	Bucket string `mapstructure:"AWS_BUCKET"`
}
