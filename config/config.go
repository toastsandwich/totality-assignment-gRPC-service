package config

import (
	"github.com/WistleBlowers/totality-assignment-RESTful-version/db"
	"github.com/WistleBlowers/totality-assignment-RESTful-version/repository"
	"github.com/spf13/viper"
)

var (
	UsersRepo *repository.UserRepo
	HOST      string
	PORT      string
)

func Load() {
	UsersRepo = &repository.UserRepo{
		DB: db.UserData,
	}
}

func ReadConfig() error {
	viper.SetConfigName("config.toml")
	viper.AddConfigPath("resource/")
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	HOST = viper.GetString("server.host")
	PORT = viper.GetString("server.port")
	return nil
}
