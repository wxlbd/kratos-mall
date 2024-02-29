package service

import (
	"context"
	"kratos-admin/internal/biz"

	pb "kratos-admin/api/cart/v1"
)

type CartService struct {
	pb.UnimplementedCartServer
	cartUseCase *biz.CartUseCase
}

func NewCartService(cartUseCase *biz.CartUseCase) *CartService {
	return &CartService{
		cartUseCase: cartUseCase,
	}
}

func (s *CartService) AddItem(ctx context.Context, req *pb.AddItemRequest) (*pb.AddItemReply, error) {
	memberId := s.getMemberId(ctx)
	if err := s.cartUseCase.AddItem(ctx, &biz.AddCartItemParam{
		MemberId:     memberId,
		ProductId:    req.GetProductId(),
		ProductSkuId: req.GetProductSkuId(),
		Count:        req.GetCount(),
	}); err != nil {
		return nil, err
	}
	return &pb.AddItemReply{}, nil
}
func (s *CartService) DeleteItem(ctx context.Context, req *pb.DeleteItemRequest) (*pb.DeleteItemReply, error) {
	memberId := s.getMemberId(ctx)
	if err := s.cartUseCase.DeleteItem(ctx, &biz.DeleteCartItemParam{
		MemberId:     memberId,
		ProductId:    req.GetProductId(),
		ProductSkuId: req.GetProductSkuId(),
	}); err != nil {
		return nil, err
	}
	return &pb.DeleteItemReply{}, nil
}
func (s *CartService) GetCartList(ctx context.Context, req *pb.GetCartListRequest) (*pb.GetCartListReply, error) {
	memberId := s.getMemberId(ctx)
	list, err := s.cartUseCase.GetCartList(ctx, &biz.GetCartListParam{
		MemberId: memberId,
	})
	if err != nil {
		return nil, err
	}
	reply := &pb.GetCartListReply{
		List: make([]*pb.CartProductSku, 0, len(list)),
	}
	//  DO->DTO
	for _, item := range list {
		reply.List = append(reply.List, &pb.CartProductSku{
			ProductSkuId: item.Id,
			Count:        item.Count,
		})
	}
	return &pb.GetCartListReply{}, nil
}
func (s *CartService) UpdateItemQuantity(ctx context.Context, req *pb.UpdateItemQuantityRequest) (*pb.UpdateItemQuantityReply, error) {
	memberId := s.getMemberId(ctx)
	if err := s.cartUseCase.UpdateItemQuantity(ctx, &biz.UpdateCartItemParam{
		MemberId:     memberId,
		ProductId:    req.GetProductId(),
		ProductSkuId: req.GetProductSkuId(),
		Count:        req.GetCount(),
	}); err != nil {
		return nil, err
	}
	return &pb.UpdateItemQuantityReply{}, nil
}
func (s *CartService) ClearCart(ctx context.Context, req *pb.ClearCartRequest) (*pb.ClearCartReply, error) {
	memberId := s.getMemberId(ctx)
	if err := s.cartUseCase.ClearCart(ctx, &biz.ClearCartParam{
		MemberId: memberId,
	}); err != nil {
		return nil, err
	}
	return &pb.ClearCartReply{}, nil
}

func (s *CartService) getMemberId(ctx context.Context) int64 {
	return ctx.Value("member_id").(int64)
}
