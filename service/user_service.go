package service

import (
	"github.com/toastsandwich/totality-assignment-GRPC-version/config"
	"github.com/toastsandwich/totality-assignment-GRPC-version/models"
)

func GetUserByID(id int) (*models.User, error) {
	userRepo := config.UsersRepo
	return userRepo.GetUserByID(id)
}

func SearchByCriteria(criteria, value string) ([]*models.User, error) {
	userRepo := config.UsersRepo
	var usrs []*models.User
	var err error
	if value == "true" {
		usrs, err = userRepo.SearchByCriteria(criteria, true)
	} else if value == "false" {
		usrs, err = userRepo.SearchByCriteria(criteria, false)
	} else {
		usrs, err = userRepo.SearchByCriteria(criteria, value)
	}
	return usrs, err
}
