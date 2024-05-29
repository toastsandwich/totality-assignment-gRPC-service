package repository

import (
	"reflect"
	"testing"

	"github.com/toastsandwich/totality-assignment-GRPC-version/db"
	"github.com/toastsandwich/totality-assignment-GRPC-version/models"
)

func TestUserRepo_GetUserByID(t *testing.T) {
	userRepo := UserRepo{DB: db.UserData}

	tests := []struct {
		name      string
		id        int
		wantUser  *models.User
		wantError bool
	}{
		{"Existing ID", 1, &models.User{ID: 1, FName: "John", LName: "Doe", City: "New York", Phone: 1234567890, Height: 5.9, Married: true}, false},
		{"Non-existing ID", 100, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUser, err := userRepo.GetUserByID(tt.id)

			if (err != nil) != tt.wantError {
				t.Errorf("GetUserByID() error = %v, wantError %v", err, tt.wantError)
				return
			}

			if !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("GetUserByID() gotUser = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}

func TestUserRepo_SearchByCriteria(t *testing.T) {
	userRepo := UserRepo{DB: db.UserData}

	tests := []struct {
		name        string
		criteria    string
		val         interface{}
		wantResults []*models.User
		wantError   bool
	}{
		{"Search by city", "City", "New York", []*models.User{
			{ID: 1, FName: "John", LName: "Doe", City: "New York", Phone: 1234567890, Height: 5.9, Married: true},
			{ID: 6, FName: "Michael", LName: "Williams", City: "New York", Phone: 5678901234, Height: 6.2, Married: false},
		}, false},
		{"Search by phone", "Phone", 1234567890, []*models.User{{ID: 1, FName: "John", LName: "Doe", City: "New York", Phone: 1234567890, Height: 5.9, Married: true}}, false},
		{"Search by non-existing criteria", "NonExistingField", "Value", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResults, err := userRepo.SearchByCriteria(tt.criteria, tt.val)

			if (err != nil) != tt.wantError {
				t.Errorf("SearchByCriteria() error = %v, wantError %v", err, tt.wantError)
				return
			}

			if !reflect.DeepEqual(gotResults, tt.wantResults) {
				t.Errorf("SearchByCriteria() gotResults = %v, want %v", gotResults, tt.wantResults)
			}
		})
	}
}
