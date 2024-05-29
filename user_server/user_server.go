package user_server

import (
	"context"

	"github.com/toastsandwich/totality-assignment-GRPC-version/pb/github.com/toastsandwich/totality-assignment-GRPC-version/pb"
	"github.com/toastsandwich/totality-assignment-GRPC-version/service"
	"github.com/toastsandwich/totality-assignment-GRPC-version/utils"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func NewUserServer() *UserServer {
	return &UserServer{}
}

func (s *UserServer) GetUserByID(ctx context.Context, req *pb.GetUserByIDReq) (*pb.GetUserByIDRes, error) {
	id := req.Id
	usr, err := service.GetUserByID(int(id))
	if err != nil {
		return nil, err
	}
	pbUsr := utils.ConvertUserStructToGRPCUserStruct(usr)
	return &pb.GetUserByIDRes{
		User: pbUsr,
	}, nil
}

func (s *UserServer) SearchByCriteria(ctx context.Context, req *pb.SearchByCriteriaReq) (*pb.SearchByCriteriaRes, error) {
	criteia := req.Criteria
	value := req.Value
	usrs, err := service.SearchByCriteria(criteia, value)
	if err != nil {
		return nil, err
	}
	var pbUsrs []*pb.User
	for _, usr := range usrs {
		pbusr := utils.ConvertUserStructToGRPCUserStruct(usr)
		pbUsrs = append(pbUsrs, pbusr)
	}
	return &pb.SearchByCriteriaRes{
		Users: pbUsrs,
	}, nil
}
