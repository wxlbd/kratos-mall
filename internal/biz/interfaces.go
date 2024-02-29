package biz

import "context"

// CartRepo 购物车仓储接口
type CartRepo interface {
	AddItem(ctx context.Context, param *AddCartItemParam) error
	DeleteItem(ctx context.Context, param *DeleteCartItemParam) error
	UpdateItemQuantity(ctx context.Context, param *UpdateCartItemParam) error
	GetCartList(ctx context.Context, param *GetCartListParam) ([]*CartItem, error)
	ClearCart(ctx context.Context, param *ClearCartParam) error
}

type Product struct {
}

type ProductSku struct {
	Id    int64   // sku id
	Price float64 // 价格
	Count int64   // 数量
	Pic   string  // 图片
}

// ProductRepo 商品仓储接口
type ProductRepo interface {
	FindProductById(ctx context.Context, id int64) (*Product, error)
	// FindProductSkuBySkuId 查询商品sku
	FindProductSkuBySkuId(ctx context.Context, skuId int64) (*ProductSku, error)
	FindProductSkusBySkuIdList(ctx context.Context, skuIds []int64) ([]*ProductSku, error)
	// FindProductSkusById 查询商品sku列表
	FindProductSkusById(ctx context.Context, id int64) ([]*ProductSku, error)
}
