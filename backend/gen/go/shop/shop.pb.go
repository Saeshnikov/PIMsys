// Версия ProtoBuf

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: shop/shop.proto

// Текущий пакет - указывает пространство имен для сервиса и сообщений. Помогает избегать конфликтов имен.

package proto

import (
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

type NewShopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Url         string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *NewShopRequest) Reset() {
	*x = NewShopRequest{}
	mi := &file_shop_shop_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewShopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewShopRequest) ProtoMessage() {}

func (x *NewShopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shop_shop_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewShopRequest.ProtoReflect.Descriptor instead.
func (*NewShopRequest) Descriptor() ([]byte, []int) {
	return file_shop_shop_proto_rawDescGZIP(), []int{0}
}

func (x *NewShopRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewShopRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NewShopRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type NewShopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NewShopResponse) Reset() {
	*x = NewShopResponse{}
	mi := &file_shop_shop_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewShopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewShopResponse) ProtoMessage() {}

func (x *NewShopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shop_shop_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewShopResponse.ProtoReflect.Descriptor instead.
func (*NewShopResponse) Descriptor() ([]byte, []int) {
	return file_shop_shop_proto_rawDescGZIP(), []int{1}
}

type AlterShopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopId   int32     `protobuf:"varint,1,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
	ShopInfo *ShopInfo `protobuf:"bytes,2,opt,name=shop_info,json=shopInfo,proto3" json:"shop_info,omitempty"`
}

func (x *AlterShopRequest) Reset() {
	*x = AlterShopRequest{}
	mi := &file_shop_shop_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlterShopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlterShopRequest) ProtoMessage() {}

func (x *AlterShopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shop_shop_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlterShopRequest.ProtoReflect.Descriptor instead.
func (*AlterShopRequest) Descriptor() ([]byte, []int) {
	return file_shop_shop_proto_rawDescGZIP(), []int{2}
}

func (x *AlterShopRequest) GetShopId() int32 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

func (x *AlterShopRequest) GetShopInfo() *ShopInfo {
	if x != nil {
		return x.ShopInfo
	}
	return nil
}

type AlterShopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AlterShopResponse) Reset() {
	*x = AlterShopResponse{}
	mi := &file_shop_shop_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlterShopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlterShopResponse) ProtoMessage() {}

func (x *AlterShopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shop_shop_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlterShopResponse.ProtoReflect.Descriptor instead.
func (*AlterShopResponse) Descriptor() ([]byte, []int) {
	return file_shop_shop_proto_rawDescGZIP(), []int{3}
}

type DeleteShopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopId int32 `protobuf:"varint,1,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
}

func (x *DeleteShopRequest) Reset() {
	*x = DeleteShopRequest{}
	mi := &file_shop_shop_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteShopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteShopRequest) ProtoMessage() {}

func (x *DeleteShopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shop_shop_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteShopRequest.ProtoReflect.Descriptor instead.
func (*DeleteShopRequest) Descriptor() ([]byte, []int) {
	return file_shop_shop_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteShopRequest) GetShopId() int32 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

type DeleteShopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteShopResponse) Reset() {
	*x = DeleteShopResponse{}
	mi := &file_shop_shop_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteShopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteShopResponse) ProtoMessage() {}

func (x *DeleteShopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shop_shop_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteShopResponse.ProtoReflect.Descriptor instead.
func (*DeleteShopResponse) Descriptor() ([]byte, []int) {
	return file_shop_shop_proto_rawDescGZIP(), []int{5}
}

type ShopInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopId      int32  `protobuf:"varint,1,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Url         string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *ShopInfo) Reset() {
	*x = ShopInfo{}
	mi := &file_shop_shop_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShopInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShopInfo) ProtoMessage() {}

func (x *ShopInfo) ProtoReflect() protoreflect.Message {
	mi := &file_shop_shop_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShopInfo.ProtoReflect.Descriptor instead.
func (*ShopInfo) Descriptor() ([]byte, []int) {
	return file_shop_shop_proto_rawDescGZIP(), []int{6}
}

func (x *ShopInfo) GetShopId() int32 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

func (x *ShopInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ShopInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ShopInfo) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ListShopsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListShopsRequest) Reset() {
	*x = ListShopsRequest{}
	mi := &file_shop_shop_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListShopsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListShopsRequest) ProtoMessage() {}

func (x *ListShopsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shop_shop_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListShopsRequest.ProtoReflect.Descriptor instead.
func (*ListShopsRequest) Descriptor() ([]byte, []int) {
	return file_shop_shop_proto_rawDescGZIP(), []int{7}
}

type ListShopsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info []*ShopInfo `protobuf:"bytes,1,rep,name=info,proto3" json:"info,omitempty"`
}

func (x *ListShopsResponse) Reset() {
	*x = ListShopsResponse{}
	mi := &file_shop_shop_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListShopsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListShopsResponse) ProtoMessage() {}

func (x *ListShopsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shop_shop_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListShopsResponse.ProtoReflect.Descriptor instead.
func (*ListShopsResponse) Descriptor() ([]byte, []int) {
	return file_shop_shop_proto_rawDescGZIP(), []int{8}
}

func (x *ListShopsResponse) GetInfo() []*ShopInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_shop_shop_proto protoreflect.FileDescriptor

var file_shop_shop_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x73, 0x68, 0x6f, 0x70, 0x22, 0x58, 0x0a, 0x0e, 0x4e, 0x65, 0x77, 0x53, 0x68,
	0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x22, 0x11, 0x0a, 0x0f, 0x4e, 0x65, 0x77, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x58, 0x0a, 0x10, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x68, 0x6f,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49,
	0x64, 0x12, 0x2b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x53, 0x68, 0x6f, 0x70,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x13,
	0x0a, 0x11, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x2c, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x6f,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49,
	0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6b, 0x0a, 0x08, 0x53, 0x68, 0x6f, 0x70, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x6f, 0x70,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x37, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x68, 0x6f, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x68,
	0x6f, 0x70, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x32, 0xfb, 0x01, 0x0a, 0x04, 0x53, 0x68, 0x6f, 0x70, 0x12, 0x3c, 0x0a, 0x09, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x73, 0x12, 0x16, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x53,
	0x68, 0x6f, 0x70, 0x12, 0x14, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x4e, 0x65, 0x77, 0x53, 0x68,
	0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x68, 0x6f, 0x70,
	0x2e, 0x4e, 0x65, 0x77, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3c, 0x0a, 0x09, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x12, 0x16, 0x2e,
	0x73, 0x68, 0x6f, 0x70, 0x2e, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x41, 0x6c, 0x74,
	0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f,
	0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x70, 0x12, 0x17, 0x2e, 0x73,
	0x68, 0x6f, 0x70, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x18, 0x5a, 0x16, 0x70, 0x69, 0x6d, 0x2d, 0x73, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x76, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_shop_shop_proto_rawDescOnce sync.Once
	file_shop_shop_proto_rawDescData = file_shop_shop_proto_rawDesc
)

func file_shop_shop_proto_rawDescGZIP() []byte {
	file_shop_shop_proto_rawDescOnce.Do(func() {
		file_shop_shop_proto_rawDescData = protoimpl.X.CompressGZIP(file_shop_shop_proto_rawDescData)
	})
	return file_shop_shop_proto_rawDescData
}

var file_shop_shop_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_shop_shop_proto_goTypes = []any{
	(*NewShopRequest)(nil),     // 0: shop.NewShopRequest
	(*NewShopResponse)(nil),    // 1: shop.NewShopResponse
	(*AlterShopRequest)(nil),   // 2: shop.AlterShopRequest
	(*AlterShopResponse)(nil),  // 3: shop.AlterShopResponse
	(*DeleteShopRequest)(nil),  // 4: shop.DeleteShopRequest
	(*DeleteShopResponse)(nil), // 5: shop.DeleteShopResponse
	(*ShopInfo)(nil),           // 6: shop.ShopInfo
	(*ListShopsRequest)(nil),   // 7: shop.ListShopsRequest
	(*ListShopsResponse)(nil),  // 8: shop.ListShopsResponse
}
var file_shop_shop_proto_depIdxs = []int32{
	6, // 0: shop.AlterShopRequest.shop_info:type_name -> shop.ShopInfo
	6, // 1: shop.ListShopsResponse.info:type_name -> shop.ShopInfo
	7, // 2: shop.Shop.ListShops:input_type -> shop.ListShopsRequest
	0, // 3: shop.Shop.NewShop:input_type -> shop.NewShopRequest
	2, // 4: shop.Shop.AlterShop:input_type -> shop.AlterShopRequest
	4, // 5: shop.Shop.DeleteShop:input_type -> shop.DeleteShopRequest
	8, // 6: shop.Shop.ListShops:output_type -> shop.ListShopsResponse
	1, // 7: shop.Shop.NewShop:output_type -> shop.NewShopResponse
	3, // 8: shop.Shop.AlterShop:output_type -> shop.AlterShopResponse
	5, // 9: shop.Shop.DeleteShop:output_type -> shop.DeleteShopResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_shop_shop_proto_init() }
func file_shop_shop_proto_init() {
	if File_shop_shop_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shop_shop_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shop_shop_proto_goTypes,
		DependencyIndexes: file_shop_shop_proto_depIdxs,
		MessageInfos:      file_shop_shop_proto_msgTypes,
	}.Build()
	File_shop_shop_proto = out.File
	file_shop_shop_proto_rawDesc = nil
	file_shop_shop_proto_goTypes = nil
	file_shop_shop_proto_depIdxs = nil
}
