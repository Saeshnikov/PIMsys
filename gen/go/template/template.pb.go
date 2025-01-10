// Версия ProtoBuf

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: template/template.proto

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

type NewTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string           `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	BranchId    int32            `protobuf:"varint,3,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	Attributes  []*AttributeInfo `protobuf:"bytes,4,rep,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *NewTemplateRequest) Reset() {
	*x = NewTemplateRequest{}
	mi := &file_template_template_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewTemplateRequest) ProtoMessage() {}

func (x *NewTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_template_template_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewTemplateRequest.ProtoReflect.Descriptor instead.
func (*NewTemplateRequest) Descriptor() ([]byte, []int) {
	return file_template_template_proto_rawDescGZIP(), []int{0}
}

func (x *NewTemplateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewTemplateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NewTemplateRequest) GetBranchId() int32 {
	if x != nil {
		return x.BranchId
	}
	return 0
}

func (x *NewTemplateRequest) GetAttributes() []*AttributeInfo {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type NewTemplateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NewTemplateResponse) Reset() {
	*x = NewTemplateResponse{}
	mi := &file_template_template_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewTemplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewTemplateResponse) ProtoMessage() {}

func (x *NewTemplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_template_template_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewTemplateResponse.ProtoReflect.Descriptor instead.
func (*NewTemplateResponse) Descriptor() ([]byte, []int) {
	return file_template_template_proto_rawDescGZIP(), []int{1}
}

type AlterTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BranchId    int32            `protobuf:"varint,1,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	Name        string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string           `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Attributes  []*AttributeInfo `protobuf:"bytes,4,rep,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *AlterTemplateRequest) Reset() {
	*x = AlterTemplateRequest{}
	mi := &file_template_template_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlterTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlterTemplateRequest) ProtoMessage() {}

func (x *AlterTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_template_template_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlterTemplateRequest.ProtoReflect.Descriptor instead.
func (*AlterTemplateRequest) Descriptor() ([]byte, []int) {
	return file_template_template_proto_rawDescGZIP(), []int{2}
}

func (x *AlterTemplateRequest) GetBranchId() int32 {
	if x != nil {
		return x.BranchId
	}
	return 0
}

func (x *AlterTemplateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AlterTemplateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AlterTemplateRequest) GetAttributes() []*AttributeInfo {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type AlterTemplateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AlterTemplateResponse) Reset() {
	*x = AlterTemplateResponse{}
	mi := &file_template_template_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlterTemplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlterTemplateResponse) ProtoMessage() {}

func (x *AlterTemplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_template_template_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlterTemplateResponse.ProtoReflect.Descriptor instead.
func (*AlterTemplateResponse) Descriptor() ([]byte, []int) {
	return file_template_template_proto_rawDescGZIP(), []int{3}
}

type DeleteTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId int32 `protobuf:"varint,1,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
}

func (x *DeleteTemplateRequest) Reset() {
	*x = DeleteTemplateRequest{}
	mi := &file_template_template_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTemplateRequest) ProtoMessage() {}

func (x *DeleteTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_template_template_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTemplateRequest.ProtoReflect.Descriptor instead.
func (*DeleteTemplateRequest) Descriptor() ([]byte, []int) {
	return file_template_template_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteTemplateRequest) GetTemplateId() int32 {
	if x != nil {
		return x.TemplateId
	}
	return 0
}

type DeleteTemplateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTemplateResponse) Reset() {
	*x = DeleteTemplateResponse{}
	mi := &file_template_template_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTemplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTemplateResponse) ProtoMessage() {}

func (x *DeleteTemplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_template_template_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTemplateResponse.ProtoReflect.Descriptor instead.
func (*DeleteTemplateResponse) Descriptor() ([]byte, []int) {
	return file_template_template_proto_rawDescGZIP(), []int{5}
}

type ListTemplatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListTemplatesRequest) Reset() {
	*x = ListTemplatesRequest{}
	mi := &file_template_template_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTemplatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTemplatesRequest) ProtoMessage() {}

func (x *ListTemplatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_template_template_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTemplatesRequest.ProtoReflect.Descriptor instead.
func (*ListTemplatesRequest) Descriptor() ([]byte, []int) {
	return file_template_template_proto_rawDescGZIP(), []int{6}
}

type ListTemplatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info []*TemplateInfo `protobuf:"bytes,1,rep,name=info,proto3" json:"info,omitempty"`
}

func (x *ListTemplatesResponse) Reset() {
	*x = ListTemplatesResponse{}
	mi := &file_template_template_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTemplatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTemplatesResponse) ProtoMessage() {}

func (x *ListTemplatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_template_template_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTemplatesResponse.ProtoReflect.Descriptor instead.
func (*ListTemplatesResponse) Descriptor() ([]byte, []int) {
	return file_template_template_proto_rawDescGZIP(), []int{7}
}

func (x *ListTemplatesResponse) GetInfo() []*TemplateInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type TemplateInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId  int32            `protobuf:"varint,1,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	Name        string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string           `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Attributes  []*AttributeInfo `protobuf:"bytes,4,rep,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *TemplateInfo) Reset() {
	*x = TemplateInfo{}
	mi := &file_template_template_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TemplateInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateInfo) ProtoMessage() {}

func (x *TemplateInfo) ProtoReflect() protoreflect.Message {
	mi := &file_template_template_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateInfo.ProtoReflect.Descriptor instead.
func (*TemplateInfo) Descriptor() ([]byte, []int) {
	return file_template_template_proto_rawDescGZIP(), []int{8}
}

func (x *TemplateInfo) GetTemplateId() int32 {
	if x != nil {
		return x.TemplateId
	}
	return 0
}

func (x *TemplateInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TemplateInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TemplateInfo) GetAttributes() []*AttributeInfo {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type AttributeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type            string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	IsValueRequired bool   `protobuf:"varint,3,opt,name=is_value_required,json=isValueRequired,proto3" json:"is_value_required,omitempty"`
	IsUnique        bool   `protobuf:"varint,4,opt,name=is_unique,json=isUnique,proto3" json:"is_unique,omitempty"`
	Name            string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Description     string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *AttributeInfo) Reset() {
	*x = AttributeInfo{}
	mi := &file_template_template_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AttributeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributeInfo) ProtoMessage() {}

func (x *AttributeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_template_template_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributeInfo.ProtoReflect.Descriptor instead.
func (*AttributeInfo) Descriptor() ([]byte, []int) {
	return file_template_template_proto_rawDescGZIP(), []int{9}
}

func (x *AttributeInfo) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AttributeInfo) GetIsValueRequired() bool {
	if x != nil {
		return x.IsValueRequired
	}
	return false
}

func (x *AttributeInfo) GetIsUnique() bool {
	if x != nil {
		return x.IsUnique
	}
	return false
}

func (x *AttributeInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AttributeInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_template_template_proto protoreflect.FileDescriptor

var file_template_template_proto_rawDesc = []byte{
	0x0a, 0x17, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x73, 0x68, 0x6f, 0x70, 0x22,
	0x9c, 0x01, 0x0a, 0x12, 0x4e, 0x65, 0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x0a, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x73, 0x68, 0x6f, 0x70, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22, 0x15,
	0x0a, 0x13, 0x4e, 0x65, 0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x9e, 0x01, 0x0a, 0x14, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x33, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22, 0x17, 0x0a, 0x15, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x38, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x22, 0x18, 0x0a, 0x16, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3f, 0x0a, 0x15, 0x4c,
	0x69, 0x73, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x9a, 0x01, 0x0a,
	0x0c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a,
	0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22, 0xa2, 0x01, 0x0a, 0x0d, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x2a, 0x0a, 0x11, 0x69, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69,
	0x73, 0x5f, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x69, 0x73, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xaf,
	0x02, 0x0a, 0x08, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x48, 0x0a, 0x0d, 0x4c,
	0x69, 0x73, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x2e, 0x73,
	0x68, 0x6f, 0x70, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x4e, 0x65, 0x77, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x4e, 0x65, 0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0d, 0x41, 0x6c, 0x74,
	0x65, 0x72, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x73, 0x68, 0x6f,
	0x70, 0x2e, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x41, 0x6c,
	0x74, 0x65, 0x72, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x18, 0x5a, 0x16, 0x70, 0x69, 0x6d, 0x2d, 0x73, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x76, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_template_template_proto_rawDescOnce sync.Once
	file_template_template_proto_rawDescData = file_template_template_proto_rawDesc
)

func file_template_template_proto_rawDescGZIP() []byte {
	file_template_template_proto_rawDescOnce.Do(func() {
		file_template_template_proto_rawDescData = protoimpl.X.CompressGZIP(file_template_template_proto_rawDescData)
	})
	return file_template_template_proto_rawDescData
}

var file_template_template_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_template_template_proto_goTypes = []any{
	(*NewTemplateRequest)(nil),     // 0: shop.NewTemplateRequest
	(*NewTemplateResponse)(nil),    // 1: shop.NewTemplateResponse
	(*AlterTemplateRequest)(nil),   // 2: shop.AlterTemplateRequest
	(*AlterTemplateResponse)(nil),  // 3: shop.AlterTemplateResponse
	(*DeleteTemplateRequest)(nil),  // 4: shop.DeleteTemplateRequest
	(*DeleteTemplateResponse)(nil), // 5: shop.DeleteTemplateResponse
	(*ListTemplatesRequest)(nil),   // 6: shop.ListTemplatesRequest
	(*ListTemplatesResponse)(nil),  // 7: shop.ListTemplatesResponse
	(*TemplateInfo)(nil),           // 8: shop.TemplateInfo
	(*AttributeInfo)(nil),          // 9: shop.AttributeInfo
}
var file_template_template_proto_depIdxs = []int32{
	9, // 0: shop.NewTemplateRequest.attributes:type_name -> shop.AttributeInfo
	9, // 1: shop.AlterTemplateRequest.attributes:type_name -> shop.AttributeInfo
	8, // 2: shop.ListTemplatesResponse.info:type_name -> shop.TemplateInfo
	9, // 3: shop.TemplateInfo.attributes:type_name -> shop.AttributeInfo
	6, // 4: shop.Template.ListTemplates:input_type -> shop.ListTemplatesRequest
	0, // 5: shop.Template.NewTemplate:input_type -> shop.NewTemplateRequest
	2, // 6: shop.Template.AlterTemplate:input_type -> shop.AlterTemplateRequest
	4, // 7: shop.Template.DeleteTemplate:input_type -> shop.DeleteTemplateRequest
	7, // 8: shop.Template.ListTemplates:output_type -> shop.ListTemplatesResponse
	1, // 9: shop.Template.NewTemplate:output_type -> shop.NewTemplateResponse
	3, // 10: shop.Template.AlterTemplate:output_type -> shop.AlterTemplateResponse
	5, // 11: shop.Template.DeleteTemplate:output_type -> shop.DeleteTemplateResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_template_template_proto_init() }
func file_template_template_proto_init() {
	if File_template_template_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_template_template_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_template_template_proto_goTypes,
		DependencyIndexes: file_template_template_proto_depIdxs,
		MessageInfos:      file_template_template_proto_msgTypes,
	}.Build()
	File_template_template_proto = out.File
	file_template_template_proto_rawDesc = nil
	file_template_template_proto_goTypes = nil
	file_template_template_proto_depIdxs = nil
}
