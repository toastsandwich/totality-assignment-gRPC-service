package service

import (
	"strconv"

	"github.com/WistleBlowers/totality-assignment-RESTful-version/config"
	"github.com/WistleBlowers/totality-assignment-RESTful-version/models"
)

func GetUserByID(id_str string) (*models.User, error) {
	userRepo := config.UsersRepo
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return nil, err
	}
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
