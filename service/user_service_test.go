package service

import (
	"reflect"
	"testing"

	"github.com/toastsandwich/totality-assignment-GRPC-version/config"
	"github.com/toastsandwich/totality-assignment-GRPC-version/models"
	"github.com/toastsandwich/totality-assignment-GRPC-version/repository"
)

func TestGetUserByID(t *testing.T) {
	config.UsersRepo = &repository.UserRepo{DB: []*models.User{
		{ID: 1, FName: "John", LName: "Doe", City: "New York", Phone: 1234567890, Height: 5.9, Married: true},
	}}

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
			gotUser, err := GetUserByID(tt.id)

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

func TestSearchByCriteria(t *testing.T) {
	config.UsersRepo = &repository.UserRepo{DB: []*models.User{
		{ID: 1, FName: "John", LName: "Doe", City: "New York", Phone: 1234567890, Height: 5.9, Married: true},
		{ID: 2, FName: "Jane", LName: "Smith", City: "Los Angeles", Phone: 2345678901, Height: 5.5, Married: false},
	}}

	tests := []struct {
		name        string
		criteria    string
		value       string
		wantResults []*models.User
		wantError   bool
	}{
		{"Search by city", "City", "New York", []*models.User{{ID: 1, FName: "John", LName: "Doe", City: "New York", Phone: 1234567890, Height: 5.9, Married: true}}, false},
		{"Search by non-existing city", "City", "Miami", []*models.User{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResults, err := SearchByCriteria(tt.criteria, tt.value)

			if (err != nil) != tt.wantError {
				t.Errorf("SearchByCriteria() error = %v, wantError %v", err, tt.wantError)
				return
			}

			if len(gotResults) != len(tt.wantResults) {
				t.Errorf("SearchByCriteria() gotResults length = %d, want %d", len(gotResults), len(tt.wantResults))
			}

			for i := range gotResults {
				if !reflect.DeepEqual(gotResults[i], tt.wantResults[i]) {
					t.Errorf("SearchByCriteria() gotResults[%d] = %v, want %v", i, gotResults[i], tt.wantResults[i])
				}
			}
		})
	}
}
