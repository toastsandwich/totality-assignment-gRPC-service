package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
	"github.com/toastsandwich/totality-assignment-GRPC-version/db"
	"github.com/toastsandwich/totality-assignment-GRPC-version/repository"
)

var (
	UsersRepo *repository.UserRepo
	HOST      string
	PORT      string
	LOGPATH   string
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
	LOGPATH = viper.GetString("log.path")
	return nil
}

func InitLogger() (*os.File, error) {
	file, err := os.OpenFile(LOGPATH, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	return file, nil
}
