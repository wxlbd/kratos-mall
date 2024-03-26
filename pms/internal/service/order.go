package service

import (
	"context"

	"github.com/wxlbd/kratos-pms/internal/biz"

	pb "github.com/wxlbd/kratos-pms/api/order/v1"
)

type OrderService struct {
	pb.UnimplementedOrderServer
	cartUseCase *biz.CartUseCase
}

func NewOrderService(cartUseCase *biz.CartUseCase) *OrderService {
	return &OrderService{
		cartUseCase: cartUseCase,
	}
}

func (s *OrderService) AddCartItem(ctx context.Context, req *pb.AddCartItemRequest) (*pb.AddCartItemReply, error) {
	req.MemberId = s.getMemberId(ctx)
	if err := s.cartUseCase.AddItem(ctx, req); err != nil {
		return nil, err
	}
	return &pb.AddCartItemReply{}, nil
}

func (s *OrderService) DeleteCartItem(ctx context.Context, req *pb.DeleteCartItemRequest) (*pb.DeleteCartItemReply, error) {
	req.MemberId = s.getMemberId(ctx)
	return &pb.DeleteCartItemReply{}, nil
}

func (s *OrderService) getMemberId(ctx context.Context) string {
	memberId := ctx.Value("memberId")
	mid := memberId.(string)
	return mid
}

func (s *OrderService) GetCartList(ctx context.Context, req *pb.GetCartListRequest) (*pb.GetCartListReply, error) {
	req.MemberId = s.getMemberId(ctx)
	items, err := s.cartUseCase.GetCartList(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.GetCartListReply{
		List: items,
	}, nil
}

func (s *OrderService) UpdateCartItemQuantity(ctx context.Context, req *pb.UpdateCartItemQuantityRequest) (*pb.UpdateCartItemQuantityReply, error) {
	req.MemberId = s.getMemberId(ctx)
	if err := s.cartUseCase.UpdateItemQuantity(ctx, req); err != nil {
		return nil, err
	}
	return &pb.UpdateCartItemQuantityReply{}, nil
}

func (s *OrderService) ClearCart(ctx context.Context, req *pb.ClearCartRequest) (*pb.ClearCartReply, error) {
	req.MemberId = s.getMemberId(ctx)
	if err := s.cartUseCase.ClearCart(ctx, req); err != nil {
		return nil, err
	}
	return &pb.ClearCartReply{}, nil
}
