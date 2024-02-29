package data

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/redis/go-redis/v9"
	v1 "kratos-admin/api"
	"kratos-admin/internal/biz"
	"strconv"
)

type CartRepo struct {
	data *Data
}

type Val map[int64]int64

func (v *Val) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, v)
}

func (v Val) MarshalBinary() (data []byte, err error) {
	return json.Marshal(v)
}

// AddItem 添加商品到购物车
func (c *CartRepo) AddItem(ctx context.Context, param *biz.AddCartItemParam) error {
	if err := c.data.HSet(ctx, fmt.Sprintf("cart:%d", param.GetMemberId()), param.ProductId, param.ProductSkuId, param.Count).Err(); err != nil {
		return v1.ErrorDbError("failed to add product to cart").WithCause(err)
	}
	return nil
}

// DeleteItem 删除购物车中的一个SKU
func (c *CartRepo) DeleteItem(ctx context.Context, param *biz.DeleteCartItemParam) error {
	// 根据用户ID和商品ID查询购物车中的商品
	var val Val
	if err := c.data.HGet(ctx, fmt.Sprintf("cart:%d", param.GetMemberId()), strconv.FormatInt(param.ProductId, 10)).Scan(&val); err != nil {
		if errors.Is(err, redis.Nil) {
			return nil
		}
		return v1.ErrorDbError("failed to delete product from cart").WithCause(err)
	}
	// 删除商品SKU
	delete(val, param.ProductSkuId)
	// 更新购物车
	err := c.data.HSet(ctx, fmt.Sprintf("cart:%d", param.GetMemberId()), param.ProductId, val).Err()
	if err != nil {
		return v1.ErrorDbError("failed to delete product from cart").WithCause(err)
	}
	return nil
}

// UpdateItemQuantity 更新购物车中的一个SKU的数量
func (c *CartRepo) UpdateItemQuantity(ctx context.Context, param *biz.UpdateCartItemParam) error {
	// 根据用户ID和商品ID查询购物车中的商品
	var val Val
	if err := c.data.HGet(ctx, fmt.Sprintf("cart:%d", param.GetMemberId()), strconv.FormatInt(param.ProductId, 10)).Scan(&val); err != nil {
		return err
	}
	// 更新商品SKU数量
	val[param.ProductSkuId] = param.Count
	// 更新购物车
	err := c.data.HSet(ctx, fmt.Sprintf("cart:%d", param.GetMemberId()), param.ProductId, val).Err()
	if err != nil {
		return err
	}
	return nil
}

type CartAll map[int64]*Val

// GetCartList 获取购物车所有商品
func (c *CartRepo) GetCartList(ctx context.Context, param *biz.GetCartListParam) ([]*biz.CartItem, error) {
	// 从redis中获取购物车所有商品
	result, err := c.data.HGetAll(ctx, fmt.Sprintf("cart:%d", param.GetMemberId())).Result()
	if err != nil {
		// 如果redis中没有该商品，则返回空
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		// 如果redis中获取商品失败，则返回错误
		return nil, v1.ErrorDbError("failed to get cart list").WithCause(err)
	}
	// 初始化购物车商品列表
	var cartItems []*biz.CartItem
	// 遍历redis中购物车商品
	for k, v := range result {
		// 初始化商品信息
		var val Val
		// 将商品信息转换成Val类型
		if err := json.Unmarshal([]byte(v), &val); err != nil {
			// 如果转换失败，则返回错误
			return nil, v1.ErrorDbError("failed to get cart list").WithCause(err)
		}
		// 将商品id转换成int64类型
		productId, _ := strconv.ParseInt(k, 10, 64)
		// 遍历商品信息
		for skuId, count := range val {
			// 将商品信息添加到购物车商品列表中
			cartItems = append(cartItems, &biz.CartItem{
				ProductId:    productId,
				ProductSkuId: skuId,
				Count:        count,
			})
		}
	}
	// 返回购物车商品列表
	return nil, nil
}

// ClearCart 清空购物车
func (c *CartRepo) ClearCart(ctx context.Context, param *biz.ClearCartParam) error {
	if err := c.data.Del(ctx, fmt.Sprintf("cart:%d", param.GetMemberId())).Err(); err != nil {
		return v1.ErrorDbError("failed to clear cart").WithCause(err)
	}
	return nil
}
