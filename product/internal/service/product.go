package service

import (
	"context"
	"xshop/product/internal/data"
	pb "xshop/product/protos"
)

type Service struct {
	pb.UnimplementedProductServer
}

func (s *Service) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateReply, error) {
	newProduct := data.Product{
		Name:   in.Name,
		Desc:   in.Desc,
		Stock:  in.Stock,
		Amount: in.Amount,
		Status: in.Status,
	}
	data.DbConn.Create(&newProduct)
	return &pb.CreateReply{Id: 1}, nil
}
