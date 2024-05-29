package config

import (
	"github.com/WistleBlowers/totality-assignment-RESTful-version/db"
	"github.com/WistleBlowers/totality-assignment-RESTful-version/repository"
)

var UsersRepo *repository.UserRepo

func Load() {
	UsersRepo = &repository.UserRepo{
		DB: db.UserData,
	}
}
