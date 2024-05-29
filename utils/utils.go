package utils

import (
	"encoding/json"
	"net/http"

	"github.com/toastsandwich/totality-assignment-GRPC-version/models"
	"github.com/toastsandwich/totality-assignment-GRPC-version/pb/github.com/toastsandwich/totality-assignment-GRPC-version/pb"
)

func CreateAndWriteJSON(w http.ResponseWriter, v any) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(v)
}

func ConvertUserStructToGRPCUserStruct(u *models.User) *pb.User {
	return &pb.User{
		Fname:   u.FName,
		Lname:   u.LName,
		City:    u.City,
		Id:      int32(u.ID),
		Phone:   int32(u.Phone),
		Height:  u.Height,
		Married: u.Married,
	}

}
