// Версия ProtoBuf

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: template/template.proto

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
	Template_ListTemplates_FullMethodName  = "/shop.Template/ListTemplates"
	Template_NewTemplate_FullMethodName    = "/shop.Template/NewTemplate"
	Template_DeleteTemplate_FullMethodName = "/shop.Template/DeleteTemplate"
)

// TemplateClient is the client API for Template service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TemplateClient interface {
	ListTemplates(ctx context.Context, in *ListTemplatesRequest, opts ...grpc.CallOption) (*ListTemplatesResponse, error)
	NewTemplate(ctx context.Context, in *NewTemplateRequest, opts ...grpc.CallOption) (*NewTemplateResponse, error)
	DeleteTemplate(ctx context.Context, in *DeleteTemplateRequest, opts ...grpc.CallOption) (*DeleteTemplateResponse, error)
}

type templateClient struct {
	cc grpc.ClientConnInterface
}

func NewTemplateClient(cc grpc.ClientConnInterface) TemplateClient {
	return &templateClient{cc}
}

func (c *templateClient) ListTemplates(ctx context.Context, in *ListTemplatesRequest, opts ...grpc.CallOption) (*ListTemplatesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListTemplatesResponse)
	err := c.cc.Invoke(ctx, Template_ListTemplates_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) NewTemplate(ctx context.Context, in *NewTemplateRequest, opts ...grpc.CallOption) (*NewTemplateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NewTemplateResponse)
	err := c.cc.Invoke(ctx, Template_NewTemplate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) DeleteTemplate(ctx context.Context, in *DeleteTemplateRequest, opts ...grpc.CallOption) (*DeleteTemplateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteTemplateResponse)
	err := c.cc.Invoke(ctx, Template_DeleteTemplate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TemplateServer is the server API for Template service.
// All implementations must embed UnimplementedTemplateServer
// for forward compatibility.
type TemplateServer interface {
	ListTemplates(context.Context, *ListTemplatesRequest) (*ListTemplatesResponse, error)
	NewTemplate(context.Context, *NewTemplateRequest) (*NewTemplateResponse, error)
	DeleteTemplate(context.Context, *DeleteTemplateRequest) (*DeleteTemplateResponse, error)
	mustEmbedUnimplementedTemplateServer()
}

// UnimplementedTemplateServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTemplateServer struct{}

func (UnimplementedTemplateServer) ListTemplates(context.Context, *ListTemplatesRequest) (*ListTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTemplates not implemented")
}
func (UnimplementedTemplateServer) NewTemplate(context.Context, *NewTemplateRequest) (*NewTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewTemplate not implemented")
}
func (UnimplementedTemplateServer) DeleteTemplate(context.Context, *DeleteTemplateRequest) (*DeleteTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTemplate not implemented")
}
func (UnimplementedTemplateServer) mustEmbedUnimplementedTemplateServer() {}
func (UnimplementedTemplateServer) testEmbeddedByValue()                  {}

// UnsafeTemplateServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TemplateServer will
// result in compilation errors.
type UnsafeTemplateServer interface {
	mustEmbedUnimplementedTemplateServer()
}

func RegisterTemplateServer(s grpc.ServiceRegistrar, srv TemplateServer) {
	// If the following call pancis, it indicates UnimplementedTemplateServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Template_ServiceDesc, srv)
}

func _Template_ListTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).ListTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Template_ListTemplates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).ListTemplates(ctx, req.(*ListTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_NewTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).NewTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Template_NewTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).NewTemplate(ctx, req.(*NewTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_DeleteTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).DeleteTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Template_DeleteTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).DeleteTemplate(ctx, req.(*DeleteTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Template_ServiceDesc is the grpc.ServiceDesc for Template service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Template_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shop.Template",
	HandlerType: (*TemplateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTemplates",
			Handler:    _Template_ListTemplates_Handler,
		},
		{
			MethodName: "NewTemplate",
			Handler:    _Template_NewTemplate_Handler,
		},
		{
			MethodName: "DeleteTemplate",
			Handler:    _Template_DeleteTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "template/template.proto",
}
