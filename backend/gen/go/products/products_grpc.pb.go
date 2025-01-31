// Версия ProtoBuf

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: products/products.proto

// Текущий пакет - указывает пространство имен для сервиса и сообщений. Помогает избегать конфликтов имен.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Product_NewProduct_FullMethodName    = "/products.Product/NewProduct"
	Product_AlterProduct_FullMethodName  = "/products.Product/AlterProduct"
	Product_DeleteProduct_FullMethodName = "/products.Product/DeleteProduct"
	Product_ListProducts_FullMethodName  = "/products.Product/ListProducts"
	Product_SellProduct_FullMethodName   = "/products.Product/SellProduct"
)

// ProductClient is the client API for Product service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductClient interface {
	NewProduct(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Empty, error)
	AlterProduct(ctx context.Context, in *ProductInfoWithId, opts ...grpc.CallOption) (*Empty, error)
	DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*Empty, error)
	ListProducts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Products, error)
	SellProduct(ctx context.Context, in *SellProductRequest, opts ...grpc.CallOption) (*Empty, error)
}

type productClient struct {
	cc grpc.ClientConnInterface
}

func NewProductClient(cc grpc.ClientConnInterface) ProductClient {
	return &productClient{cc}
}

func (c *productClient) NewProduct(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Product_NewProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) AlterProduct(ctx context.Context, in *ProductInfoWithId, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Product_AlterProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Product_DeleteProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) ListProducts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Products, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Products)
	err := c.cc.Invoke(ctx, Product_ListProducts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) SellProduct(ctx context.Context, in *SellProductRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Product_SellProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServer is the server API for Product service.
// All implementations must embed UnimplementedProductServer
// for forward compatibility.
type ProductServer interface {
	NewProduct(context.Context, *ProductInfo) (*Empty, error)
	AlterProduct(context.Context, *ProductInfoWithId) (*Empty, error)
	DeleteProduct(context.Context, *DeleteProductRequest) (*Empty, error)
	ListProducts(context.Context, *Empty) (*Products, error)
	SellProduct(context.Context, *SellProductRequest) (*Empty, error)
	mustEmbedUnimplementedProductServer()
}

// UnimplementedProductServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProductServer struct{}

func (UnimplementedProductServer) NewProduct(context.Context, *ProductInfo) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewProduct not implemented")
}
func (UnimplementedProductServer) AlterProduct(context.Context, *ProductInfoWithId) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlterProduct not implemented")
}
func (UnimplementedProductServer) DeleteProduct(context.Context, *DeleteProductRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedProductServer) ListProducts(context.Context, *Empty) (*Products, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProducts not implemented")
}
func (UnimplementedProductServer) SellProduct(context.Context, *SellProductRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SellProduct not implemented")
}
func (UnimplementedProductServer) mustEmbedUnimplementedProductServer() {}
func (UnimplementedProductServer) testEmbeddedByValue()                 {}

// UnsafeProductServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServer will
// result in compilation errors.
type UnsafeProductServer interface {
	mustEmbedUnimplementedProductServer()
}

func RegisterProductServer(s grpc.ServiceRegistrar, srv ProductServer) {
	// If the following call pancis, it indicates UnimplementedProductServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Product_ServiceDesc, srv)
}

func _Product_NewProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).NewProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_NewProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).NewProduct(ctx, req.(*ProductInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_AlterProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductInfoWithId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).AlterProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_AlterProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).AlterProduct(ctx, req.(*ProductInfoWithId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_DeleteProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).DeleteProduct(ctx, req.(*DeleteProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_ListProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).ListProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_ListProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).ListProducts(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_SellProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SellProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).SellProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_SellProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).SellProduct(ctx, req.(*SellProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Product_ServiceDesc is the grpc.ServiceDesc for Product service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Product_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "products.Product",
	HandlerType: (*ProductServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewProduct",
			Handler:    _Product_NewProduct_Handler,
		},
		{
			MethodName: "AlterProduct",
			Handler:    _Product_AlterProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _Product_DeleteProduct_Handler,
		},
		{
			MethodName: "ListProducts",
			Handler:    _Product_ListProducts_Handler,
		},
		{
			MethodName: "SellProduct",
			Handler:    _Product_SellProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "products/products.proto",
}
