package config

import (
	"github.com/spf13/viper"
)

const pathImageMenu = "./public/images/menu"

type Config struct {
	Db_driver     string `mapstructure:"DB_DRIVER"`
	Db_source     string `mapstructure:"PG_SOURCE"`
	Db_host       string `mapstructure:"DB_HOST"`
	Db_port       string `mapstructure:"DB_PORT"`
	Db_name       string `mapstructure:"DB_NAME"`
	Db_username   string `mapstructure:"DB_USERNAME"`
	Db_password   string `mapstructure:"DB_PASSWORD"`
	Db_sslmode    string `mapstructure:"DB_SSLMODE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	Api_key       string `mapstructure:"API_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func GetPathImageMenu() string {
	return pathImageMenu
}
