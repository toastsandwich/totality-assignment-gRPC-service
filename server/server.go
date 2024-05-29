package server

import (
	"log"
	"net"

	"github.com/toastsandwich/totality-assignment-GRPC-version/pb/github.com/toastsandwich/totality-assignment-GRPC-version/pb"
	"github.com/toastsandwich/totality-assignment-GRPC-version/user_server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Start(host, port string) error {
	addr := host + ":" + port
	server := grpc.NewServer()
	userServer := user_server.NewUserServer()
	reflection.Register(server)
	pb.RegisterUserServiceServer(server, userServer)

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	log.Println("grpc server hosted on ", addr)
	return server.Serve(ln)
}
