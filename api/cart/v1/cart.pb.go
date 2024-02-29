// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: api/cart/v1/cart.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CartProductSku struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductSkuId int64 `protobuf:"varint,1,opt,name=product_sku_id,json=productSkuId,proto3" json:"product_sku_id,omitempty"` // sku id
	Count        int64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`                                     // 数量
}

func (x *CartProductSku) Reset() {
	*x = CartProductSku{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_v1_cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartProductSku) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartProductSku) ProtoMessage() {}

func (x *CartProductSku) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_v1_cart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartProductSku.ProtoReflect.Descriptor instead.
func (*CartProductSku) Descriptor() ([]byte, []int) {
	return file_api_cart_v1_cart_proto_rawDescGZIP(), []int{0}
}

func (x *CartProductSku) GetProductSkuId() int64 {
	if x != nil {
		return x.ProductSkuId
	}
	return 0
}

func (x *CartProductSku) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type AddItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId    int64 `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`            // 商品id
	ProductSkuId int64 `protobuf:"varint,2,opt,name=product_sku_id,json=productSkuId,proto3" json:"product_sku_id,omitempty"` // 商品sku id
	ShopId       int64 `protobuf:"varint,3,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`                     // 店铺id
	Count        int64 `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *AddItemRequest) Reset() {
	*x = AddItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_v1_cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItemRequest) ProtoMessage() {}

func (x *AddItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_v1_cart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddItemRequest.ProtoReflect.Descriptor instead.
func (*AddItemRequest) Descriptor() ([]byte, []int) {
	return file_api_cart_v1_cart_proto_rawDescGZIP(), []int{1}
}

func (x *AddItemRequest) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *AddItemRequest) GetProductSkuId() int64 {
	if x != nil {
		return x.ProductSkuId
	}
	return 0
}

func (x *AddItemRequest) GetShopId() int64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

func (x *AddItemRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type AddItemReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddItemReply) Reset() {
	*x = AddItemReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_v1_cart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddItemReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItemReply) ProtoMessage() {}

func (x *AddItemReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_v1_cart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddItemReply.ProtoReflect.Descriptor instead.
func (*AddItemReply) Descriptor() ([]byte, []int) {
	return file_api_cart_v1_cart_proto_rawDescGZIP(), []int{2}
}

type DeleteItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId    int64 `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	ProductSkuId int64 `protobuf:"varint,2,opt,name=product_sku_id,json=productSkuId,proto3" json:"product_sku_id,omitempty"`
}

func (x *DeleteItemRequest) Reset() {
	*x = DeleteItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_v1_cart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteItemRequest) ProtoMessage() {}

func (x *DeleteItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_v1_cart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteItemRequest.ProtoReflect.Descriptor instead.
func (*DeleteItemRequest) Descriptor() ([]byte, []int) {
	return file_api_cart_v1_cart_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteItemRequest) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *DeleteItemRequest) GetProductSkuId() int64 {
	if x != nil {
		return x.ProductSkuId
	}
	return 0
}

type DeleteItemReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteItemReply) Reset() {
	*x = DeleteItemReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_v1_cart_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteItemReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteItemReply) ProtoMessage() {}

func (x *DeleteItemReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_v1_cart_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteItemReply.ProtoReflect.Descriptor instead.
func (*DeleteItemReply) Descriptor() ([]byte, []int) {
	return file_api_cart_v1_cart_proto_rawDescGZIP(), []int{4}
}

type GetCartListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberId int64 `protobuf:"varint,1,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
}

func (x *GetCartListRequest) Reset() {
	*x = GetCartListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_v1_cart_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartListRequest) ProtoMessage() {}

func (x *GetCartListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_v1_cart_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartListRequest.ProtoReflect.Descriptor instead.
func (*GetCartListRequest) Descriptor() ([]byte, []int) {
	return file_api_cart_v1_cart_proto_rawDescGZIP(), []int{5}
}

func (x *GetCartListRequest) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

type GetCartListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*CartProductSku `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *GetCartListReply) Reset() {
	*x = GetCartListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_v1_cart_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartListReply) ProtoMessage() {}

func (x *GetCartListReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_v1_cart_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartListReply.ProtoReflect.Descriptor instead.
func (*GetCartListReply) Descriptor() ([]byte, []int) {
	return file_api_cart_v1_cart_proto_rawDescGZIP(), []int{6}
}

func (x *GetCartListReply) GetList() []*CartProductSku {
	if x != nil {
		return x.List
	}
	return nil
}

type UpdateItemQuantityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId    int64 `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	ProductSkuId int64 `protobuf:"varint,2,opt,name=product_sku_id,json=productSkuId,proto3" json:"product_sku_id,omitempty"`
	Count        int64 `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *UpdateItemQuantityRequest) Reset() {
	*x = UpdateItemQuantityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_v1_cart_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateItemQuantityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateItemQuantityRequest) ProtoMessage() {}

func (x *UpdateItemQuantityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_v1_cart_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateItemQuantityRequest.ProtoReflect.Descriptor instead.
func (*UpdateItemQuantityRequest) Descriptor() ([]byte, []int) {
	return file_api_cart_v1_cart_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateItemQuantityRequest) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *UpdateItemQuantityRequest) GetProductSkuId() int64 {
	if x != nil {
		return x.ProductSkuId
	}
	return 0
}

func (x *UpdateItemQuantityRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type UpdateItemQuantityReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateItemQuantityReply) Reset() {
	*x = UpdateItemQuantityReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_v1_cart_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateItemQuantityReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateItemQuantityReply) ProtoMessage() {}

func (x *UpdateItemQuantityReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_v1_cart_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateItemQuantityReply.ProtoReflect.Descriptor instead.
func (*UpdateItemQuantityReply) Descriptor() ([]byte, []int) {
	return file_api_cart_v1_cart_proto_rawDescGZIP(), []int{8}
}

type ClearCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClearCartRequest) Reset() {
	*x = ClearCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_v1_cart_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearCartRequest) ProtoMessage() {}

func (x *ClearCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_v1_cart_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearCartRequest.ProtoReflect.Descriptor instead.
func (*ClearCartRequest) Descriptor() ([]byte, []int) {
	return file_api_cart_v1_cart_proto_rawDescGZIP(), []int{9}
}

type ClearCartReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClearCartReply) Reset() {
	*x = ClearCartReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cart_v1_cart_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearCartReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearCartReply) ProtoMessage() {}

func (x *ClearCartReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_v1_cart_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearCartReply.ProtoReflect.Descriptor instead.
func (*ClearCartReply) Descriptor() ([]byte, []int) {
	return file_api_cart_v1_cart_proto_rawDescGZIP(), []int{10}
}

var File_api_cart_v1_cart_proto protoreflect.FileDescriptor

var file_api_cart_v1_cart_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61,
	0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61,
	0x72, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x0e,
	0x43, 0x61, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x6b, 0x75, 0x12, 0x24,
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53,
	0x6b, 0x75, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x96, 0x01, 0x0a, 0x0e, 0x41,
	0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53,
	0x6b, 0x75, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x6a, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x2d, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x6b, 0x75, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20,
	0x00, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x6b, 0x75, 0x49, 0x64, 0x22,
	0x11, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x3a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x22, 0x02, 0x20, 0x00, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x22, 0x43,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x2f, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x61, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x6b, 0x75, 0x52, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x22, 0x88, 0x01, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x26, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x0e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x53, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x19,
	0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x51, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x12, 0x0a, 0x10, 0x43, 0x6c, 0x65,
	0x61, 0x72, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a,
	0x0e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32,
	0xa9, 0x04, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x74, 0x12, 0x55, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x12, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0c, 0x3a, 0x01, 0x2a, 0x22, 0x07, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x12,
	0x79, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2d, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x27, 0x2a, 0x25, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x7b, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x5e, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x43, 0x61, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x09, 0x12, 0x07, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x12, 0x94, 0x01, 0x0a, 0x12, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x3a, 0x01, 0x2a, 0x1a, 0x25, 0x76, 0x31, 0x2f, 0x63,
	0x61, 0x72, 0x74, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x7d,
	0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64,
	0x7d, 0x12, 0x58, 0x0a, 0x09, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x43, 0x61, 0x72, 0x74, 0x12, 0x1d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x65,
	0x61, 0x72, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x65, 0x61,
	0x72, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x09, 0x2a, 0x07, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x42, 0x2c, 0x0a, 0x0b, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x1b, 0x6b, 0x72,
	0x61, 0x74, 0x6f, 0x73, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x61, 0x72, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_cart_v1_cart_proto_rawDescOnce sync.Once
	file_api_cart_v1_cart_proto_rawDescData = file_api_cart_v1_cart_proto_rawDesc
)

func file_api_cart_v1_cart_proto_rawDescGZIP() []byte {
	file_api_cart_v1_cart_proto_rawDescOnce.Do(func() {
		file_api_cart_v1_cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_cart_v1_cart_proto_rawDescData)
	})
	return file_api_cart_v1_cart_proto_rawDescData
}

var file_api_cart_v1_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_cart_v1_cart_proto_goTypes = []interface{}{
	(*CartProductSku)(nil),            // 0: api.cart.v1.CartProductSku
	(*AddItemRequest)(nil),            // 1: api.cart.v1.AddItemRequest
	(*AddItemReply)(nil),              // 2: api.cart.v1.AddItemReply
	(*DeleteItemRequest)(nil),         // 3: api.cart.v1.DeleteItemRequest
	(*DeleteItemReply)(nil),           // 4: api.cart.v1.DeleteItemReply
	(*GetCartListRequest)(nil),        // 5: api.cart.v1.GetCartListRequest
	(*GetCartListReply)(nil),          // 6: api.cart.v1.GetCartListReply
	(*UpdateItemQuantityRequest)(nil), // 7: api.cart.v1.UpdateItemQuantityRequest
	(*UpdateItemQuantityReply)(nil),   // 8: api.cart.v1.UpdateItemQuantityReply
	(*ClearCartRequest)(nil),          // 9: api.cart.v1.ClearCartRequest
	(*ClearCartReply)(nil),            // 10: api.cart.v1.ClearCartReply
}
var file_api_cart_v1_cart_proto_depIdxs = []int32{
	0,  // 0: api.cart.v1.GetCartListReply.list:type_name -> api.cart.v1.CartProductSku
	1,  // 1: api.cart.v1.Cart.AddItem:input_type -> api.cart.v1.AddItemRequest
	3,  // 2: api.cart.v1.Cart.DeleteItem:input_type -> api.cart.v1.DeleteItemRequest
	5,  // 3: api.cart.v1.Cart.GetCartList:input_type -> api.cart.v1.GetCartListRequest
	7,  // 4: api.cart.v1.Cart.UpdateItemQuantity:input_type -> api.cart.v1.UpdateItemQuantityRequest
	9,  // 5: api.cart.v1.Cart.ClearCart:input_type -> api.cart.v1.ClearCartRequest
	2,  // 6: api.cart.v1.Cart.AddItem:output_type -> api.cart.v1.AddItemReply
	4,  // 7: api.cart.v1.Cart.DeleteItem:output_type -> api.cart.v1.DeleteItemReply
	6,  // 8: api.cart.v1.Cart.GetCartList:output_type -> api.cart.v1.GetCartListReply
	8,  // 9: api.cart.v1.Cart.UpdateItemQuantity:output_type -> api.cart.v1.UpdateItemQuantityReply
	10, // 10: api.cart.v1.Cart.ClearCart:output_type -> api.cart.v1.ClearCartReply
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_api_cart_v1_cart_proto_init() }
func file_api_cart_v1_cart_proto_init() {
	if File_api_cart_v1_cart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_cart_v1_cart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartProductSku); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_cart_v1_cart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddItemRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_cart_v1_cart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddItemReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_cart_v1_cart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteItemRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_cart_v1_cart_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteItemReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_cart_v1_cart_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_cart_v1_cart_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartListReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_cart_v1_cart_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateItemQuantityRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_cart_v1_cart_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateItemQuantityReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_cart_v1_cart_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearCartRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_cart_v1_cart_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearCartReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_cart_v1_cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_cart_v1_cart_proto_goTypes,
		DependencyIndexes: file_api_cart_v1_cart_proto_depIdxs,
		MessageInfos:      file_api_cart_v1_cart_proto_msgTypes,
	}.Build()
	File_api_cart_v1_cart_proto = out.File
	file_api_cart_v1_cart_proto_rawDesc = nil
	file_api_cart_v1_cart_proto_goTypes = nil
	file_api_cart_v1_cart_proto_depIdxs = nil
}
