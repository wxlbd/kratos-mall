package service

import (
	"context"

	pb "kratos-admin/api/cart/v1"
)

type CartService struct {
	pb.UnimplementedCartServer
}

func NewCartService() *CartService {
	return &CartService{}
}

func (s *CartService) AddItem(ctx context.Context, req *pb.AddItemRequest) (*pb.AddItemReply, error) {
	return &pb.AddItemReply{}, nil
}
func (s *CartService) DeleteItem(ctx context.Context, req *pb.DeleteItemRequest) (*pb.DeleteItemReply, error) {
	return &pb.DeleteItemReply{}, nil
}
func (s *CartService) GetCartList(ctx context.Context, req *pb.GetCartListRequest) (*pb.GetCartListReply, error) {
	return &pb.GetCartListReply{}, nil
}
func (s *CartService) UpdateItemQuantity(ctx context.Context, req *pb.UpdateItemQuantityRequest) (*pb.UpdateItemQuantityReply, error) {
	return &pb.UpdateItemQuantityReply{}, nil
}
func (s *CartService) ClearCart(ctx context.Context, req *pb.ClearCartRequest) (*pb.ClearCartReply, error) {
	return &pb.ClearCartReply{}, nil
}
