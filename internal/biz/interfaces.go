package biz

import "context"

type CartRepo interface {
	AddItem(ctx context.Context, param *AddCartItemParam) error
	DeleteItem(ctx context.Context, param *DeleteCartItemParam) error
	UpdateItemQuantity(ctx context.Context, param *UpdateCartItemParam) error
	GetCartList(ctx context.Context, param *GetCartListParam) ([]*CartItem, error)
	ClearCart(ctx context.Context, param *ClearCartParam) error
}
