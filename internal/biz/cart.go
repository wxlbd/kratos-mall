package biz

import (
	"context"

	v1 "kratos-admin/api/order/v1"
)

type CartUseCase struct {
	cartRepo CartRepo
}

func NewCartUseCase(cartRepo CartRepo) *CartUseCase {
	return &CartUseCase{cartRepo: cartRepo}
}

func (u *CartUseCase) AddItem(ctx context.Context, req *v1.AddCartItemRequest) error {
	// TODO: 其它逻辑如限制购买数量等
	return u.cartRepo.AddItem(ctx, req)
}

func (u *CartUseCase) DeleteItem(ctx context.Context, memberId string, productSkuIds ...string) error {
	return u.cartRepo.DeleteItem(ctx, memberId, productSkuIds...)
}

func (u *CartUseCase) UpdateItemQuantity(ctx context.Context, req *v1.UpdateCartItemQuantityRequest) error {
	return u.cartRepo.UpdateItemQuantity(ctx, req)
}

// GetCartList 获取购物车商品sku列表
func (u *CartUseCase) GetCartList(ctx context.Context, req *v1.GetCartListRequest) ([]*v1.CartItem, error) {
	items, err := u.cartRepo.FindAll(ctx, req)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (u *CartUseCase) ClearCart(ctx context.Context, req *v1.ClearCartRequest) error {
	return u.cartRepo.Clear(ctx, req)
}
