package server

import (
	"google.golang.org/grpc"

	"xshop/product/internal/service"
	pb "xshop/product/protos"
)

func NewGRPCServer() *grpc.Server {
	var opts = []grpc.ServerOption{}

	srv := grpc.NewServer(opts...)

	pb.RegisterProductServer(srv, &service.Service{})

	return srv
}
