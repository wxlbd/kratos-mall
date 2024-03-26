package data

import (
	"context"
	"encoding/json"
	"fmt"

	v1 "github.com/wxlbd/kratos-pms/api/order/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/redis/go-redis/v9"
	api "github.com/wxlbd/kratos-pms/api"
	"github.com/wxlbd/kratos-pms/internal/biz"
)

const (
	CartKey = "member_cart:%s"
)

type CartRepo struct {
	data *Data
}

func NewCartRepo(data *Data) biz.CartRepo {
	return &CartRepo{data: data}
}

// Item 购物车中的sku
type Item struct {
	Count uint32  `json:"count"`
	Price float64 `json:"price"`
}

func (v *Item) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, v)
}

func (v Item) MarshalBinary() (data []byte, err error) {
	return json.Marshal(v)
}

// AddItem 添加商品到购物车
func (c *CartRepo) AddItem(ctx context.Context, req *v1.AddCartItemRequest) error {
	if err := c.data.HSet(ctx, fmt.Sprintf(CartKey, req.GetMemberId()), req.ProductSkuId, &Item{Count: req.Count, Price: req.Price}).Err(); err != nil {
		return api.ErrorDbError("Failed to add item to cart").WithCause(err)
	}
	return nil
}

// DeleteItem 删除购物车中的一个SKU
func (c *CartRepo) DeleteItem(ctx context.Context, memberId string, productSkuIds ...string) error {
	key := fmt.Sprintf(CartKey, memberId)
	err := c.data.HDel(ctx, key, productSkuIds...).Err()
	if err != nil {
		return api.ErrorDbError("Failed to delete item from cart").WithCause(err)
	}
	return nil
}

// UpdateItemQuantity 更新购物车中的一个SKU的数量
func (c *CartRepo) UpdateItemQuantity(ctx context.Context, req *v1.UpdateCartItemQuantityRequest) error {
	key := fmt.Sprintf(CartKey, req.GetMemberId())
	// 根据用户ID和商品ID查询购物车中的商品
	var val Item
	if err := c.data.HGet(ctx, key, req.GetProductSkuId()).Scan(&val); err != nil {
		return err
	}
	// 更新商品SKU数量
	val.Count = req.Count
	// 更新购物车
	err := c.data.HSet(ctx, key, req.ProductSkuId, val).Err()
	if err != nil {
		return err
	}
	return nil
}

type CartAll map[int64]*Item

// FindAll 获取购物车所有商品
func (c *CartRepo) FindAll(ctx context.Context, req *v1.GetCartListRequest) ([]*v1.CartItem, error) {
	// 从redis中获取购物车所有商品
	result, err := c.data.HGetAll(ctx, fmt.Sprintf(CartKey, req.GetMemberId())).Result()
	if err != nil {
		// 如果redis中没有该商品，则返回空
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		// 如果redis中获取商品失败，则返回错误
		return nil, api.ErrorDbError("Failed to get cart item list").WithCause(err)
	}
	// 初始化购物车商品列表
	var cartItems []*v1.CartItem
	// 遍历redis中购物车商品
	for k, v := range result {
		// 初始化商品信息
		var val Item
		// 将商品信息转换成Val类型
		if err := json.Unmarshal([]byte(v), &val); err != nil {
			// 如果转换失败，则返回错误
			return nil, api.ErrorDbError("Failed to get cart item list").WithCause(err)
		}
		cartItems = append(cartItems, &v1.CartItem{
			ProductSkuId: k,
			Count:        val.Count,
			Price:        val.Price,
		})
	}
	// 返回购物车商品列表
	return cartItems, nil
}

// Clear 清空购物车
func (c *CartRepo) Clear(ctx context.Context, req *v1.ClearCartRequest) error {
	if err := c.data.Del(ctx, fmt.Sprintf(CartKey, req.GetMemberId())).Err(); err != nil {
		return api.ErrorDbError("Failed to clear cart").WithCause(err)
	}
	return nil
}
