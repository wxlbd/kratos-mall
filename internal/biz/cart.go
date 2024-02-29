package biz

import (
	"context"
)

type CartUseCase struct {
	cartRepo    CartRepo
	productRepo ProductRepo
}

func NewCartUseCase(cartRepo CartRepo, productRepo ProductRepo) *CartUseCase {
	return &CartUseCase{cartRepo: cartRepo, productRepo: productRepo}
}

func (u *CartUseCase) AddItem(ctx context.Context, req *AddCartItemParam) error {
	// TODO: 其它逻辑如限制购买数量等
	return u.cartRepo.AddItem(ctx, req)
}

func (u *CartUseCase) DeleteItem(ctx context.Context, param *DeleteCartItemParam) error {
	return u.cartRepo.DeleteItem(ctx, param)
}

func (u *CartUseCase) UpdateItemQuantity(ctx context.Context, param *UpdateCartItemParam) error {
	return u.cartRepo.UpdateItemQuantity(ctx, param)
}

// GetCartList 获取购物车商品sku列表
func (u *CartUseCase) GetCartList(ctx context.Context, param *GetCartListParam) ([]*ProductSku, error) {
	// 从购物车仓库获取购物车商品列表
	list, err := u.cartRepo.GetCartList(ctx, param)
	if err != nil {
		return nil, err
	}
	// 获取购物车商品列表长度
	length := len(list)
	// 创建一个ProductSku切片，长度为购物车商品列表长度
	productSkus := make([]*ProductSku, 0, length)
	// 创建一个int64类型的切片，长度为购物车商品列表长度
	skuIds := make([]int64, 0, length)
	// 创建一个map，key为int64类型的skuId，value为int64类型的商品数量
	skuCountMap := make(map[int64]int64, length)
	// 遍历购物车商品列表，将skuId和商品数量添加到skuIds和skuCountMap中
	for _, item := range list {
		skuIds = append(skuIds, item.ProductSkuId)
		skuCountMap[item.ProductSkuId] = item.Count
	}
	// 从商品仓库根据skuId列表获取商品sku列表
	skus, err := u.productRepo.FindProductSkusBySkuIdList(ctx, skuIds)
	if err != nil {
		return nil, err
	}
	// 遍历商品sku列表，将商品sku的价格、数量、图片添加到productSkus中
	for _, sku := range skus {
		productSkus = append(productSkus, &ProductSku{
			Price: sku.Price,
			Count: skuCountMap[sku.Id],
			Pic:   sku.Pic,
		})
	}
	// 返回商品sku列表
	return productSkus, nil
}

func (u *CartUseCase) ClearCart(ctx context.Context, param *ClearCartParam) error {
	return u.cartRepo.ClearCart(ctx, param)
}
