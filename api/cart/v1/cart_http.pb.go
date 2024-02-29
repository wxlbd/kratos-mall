// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.2
// - protoc             v4.25.3
// source: api/cart/v1/cart.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationCartAddItem = "/api.cart.v1.Cart/AddItem"
const OperationCartClearCart = "/api.cart.v1.Cart/ClearCart"
const OperationCartDeleteItem = "/api.cart.v1.Cart/DeleteItem"
const OperationCartGetCartList = "/api.cart.v1.Cart/GetCartList"
const OperationCartUpdateItemQuantity = "/api.cart.v1.Cart/UpdateItemQuantity"

type CartHTTPServer interface {
	// AddItem 添加sku到购物车
	AddItem(context.Context, *AddItemRequest) (*AddItemReply, error)
	// ClearCart 清空购物车
	ClearCart(context.Context, *ClearCartRequest) (*ClearCartReply, error)
	// DeleteItem 删除购物车中sku
	DeleteItem(context.Context, *DeleteItemRequest) (*DeleteItemReply, error)
	// GetCartList 获取购物车商品列表
	GetCartList(context.Context, *GetCartListRequest) (*GetCartListReply, error)
	// UpdateItemQuantity 更新购物车商品数量
	UpdateItemQuantity(context.Context, *UpdateItemQuantityRequest) (*UpdateItemQuantityReply, error)
}

func RegisterCartHTTPServer(s *http.Server, srv CartHTTPServer) {
	r := s.Route("/")
	r.POST("v1/cart", _Cart_AddItem0_HTTP_Handler(srv))
	r.DELETE("v1/cart/{product_id}/{product_sku_id}", _Cart_DeleteItem0_HTTP_Handler(srv))
	r.GET("v1/cart", _Cart_GetCartList0_HTTP_Handler(srv))
	r.PUT("v1/cart/{product_id}/{product_sku_id}", _Cart_UpdateItemQuantity0_HTTP_Handler(srv))
	r.DELETE("v1/cart", _Cart_ClearCart0_HTTP_Handler(srv))
}

func _Cart_AddItem0_HTTP_Handler(srv CartHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddItemRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCartAddItem)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddItem(ctx, req.(*AddItemRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AddItemReply)
		return ctx.Result(200, reply)
	}
}

func _Cart_DeleteItem0_HTTP_Handler(srv CartHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteItemRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCartDeleteItem)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteItem(ctx, req.(*DeleteItemRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteItemReply)
		return ctx.Result(200, reply)
	}
}

func _Cart_GetCartList0_HTTP_Handler(srv CartHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCartListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCartGetCartList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCartList(ctx, req.(*GetCartListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCartListReply)
		return ctx.Result(200, reply)
	}
}

func _Cart_UpdateItemQuantity0_HTTP_Handler(srv CartHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateItemQuantityRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCartUpdateItemQuantity)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateItemQuantity(ctx, req.(*UpdateItemQuantityRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateItemQuantityReply)
		return ctx.Result(200, reply)
	}
}

func _Cart_ClearCart0_HTTP_Handler(srv CartHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ClearCartRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCartClearCart)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ClearCart(ctx, req.(*ClearCartRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ClearCartReply)
		return ctx.Result(200, reply)
	}
}

type CartHTTPClient interface {
	AddItem(ctx context.Context, req *AddItemRequest, opts ...http.CallOption) (rsp *AddItemReply, err error)
	ClearCart(ctx context.Context, req *ClearCartRequest, opts ...http.CallOption) (rsp *ClearCartReply, err error)
	DeleteItem(ctx context.Context, req *DeleteItemRequest, opts ...http.CallOption) (rsp *DeleteItemReply, err error)
	GetCartList(ctx context.Context, req *GetCartListRequest, opts ...http.CallOption) (rsp *GetCartListReply, err error)
	UpdateItemQuantity(ctx context.Context, req *UpdateItemQuantityRequest, opts ...http.CallOption) (rsp *UpdateItemQuantityReply, err error)
}

type CartHTTPClientImpl struct {
	cc *http.Client
}

func NewCartHTTPClient(client *http.Client) CartHTTPClient {
	return &CartHTTPClientImpl{client}
}

func (c *CartHTTPClientImpl) AddItem(ctx context.Context, in *AddItemRequest, opts ...http.CallOption) (*AddItemReply, error) {
	var out AddItemReply
	pattern := "v1/cart"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCartAddItem))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CartHTTPClientImpl) ClearCart(ctx context.Context, in *ClearCartRequest, opts ...http.CallOption) (*ClearCartReply, error) {
	var out ClearCartReply
	pattern := "v1/cart"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCartClearCart))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CartHTTPClientImpl) DeleteItem(ctx context.Context, in *DeleteItemRequest, opts ...http.CallOption) (*DeleteItemReply, error) {
	var out DeleteItemReply
	pattern := "v1/cart/{product_id}/{product_sku_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCartDeleteItem))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CartHTTPClientImpl) GetCartList(ctx context.Context, in *GetCartListRequest, opts ...http.CallOption) (*GetCartListReply, error) {
	var out GetCartListReply
	pattern := "v1/cart"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCartGetCartList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CartHTTPClientImpl) UpdateItemQuantity(ctx context.Context, in *UpdateItemQuantityRequest, opts ...http.CallOption) (*UpdateItemQuantityReply, error) {
	var out UpdateItemQuantityReply
	pattern := "v1/cart/{product_id}/{product_sku_id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCartUpdateItemQuantity))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}