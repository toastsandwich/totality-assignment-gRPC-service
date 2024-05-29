package repository

import (
	"fmt"
	"reflect"

	"github.com/WistleBlowers/totality-assignment-RESTful-version/models"
)

type UserRepo struct {
	DB []*models.User
}

func (u *UserRepo) GetUserByID(id int) (*models.User, error) {
	for _, val := range u.DB {
		if val.ID == id {
			return val, nil
		}
	}
	return nil, fmt.Errorf("record not found")
}

func (u *UserRepo) SearchByCriteria(criteria string, val any) ([]*models.User, error) {
	v := reflect.ValueOf(u.DB)
	var result []*models.User
	for i := 0; i < v.Len(); i++ {
		item := v.Index(i).Interface()
		itemV := reflect.ValueOf(item)
		if itemV.Kind() == reflect.Ptr {
			itemV = itemV.Elem()
		}

		if itemV.Kind() != reflect.Struct {
			continue
		}

		fieldVal := itemV.FieldByName(criteria)
		if !fieldVal.IsValid() {
			return nil, fmt.Errorf("no such field: %s in item", criteria)
		}

		if reflect.DeepEqual(fieldVal.Interface(), val) {
			result = append(result, item.(*models.User))
		}
	}
	return result, nil
}
