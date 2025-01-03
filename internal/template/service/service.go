package template_service

import (
	"context"
	"fmt"
	proto "pim-sys/gen/go/template"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServerAPI struct {
	proto.UnimplementedTemplateServer // Хитрая штука, о ней ниже
	template                          Template
}

// Тот самый интерфейс, котрый мы передавали в grpcApp
type Template interface {
	NewTemplate(
		ctx context.Context,
		name string,
		description string,
		attribute_id []int32,
	) error
	AlterTemplate(
		ctx context.Context,
		name string,
		description string,
		attribute_id []int32,
	) error
	DeleteTemplate(
		ctx context.Context,
		template_id int32,
	) error
	ListTemplates(
		ctx context.Context,
	) (
		[]*proto.TemplateInfo,
		error,
	)
}

func Register(gRPCServer *grpc.Server, template Template) {
	proto.RegisterTemplateServer(gRPCServer, &ServerAPI{template: template})
}

func (s *ServerAPI) NewTemplate(
	ctx context.Context,
	in *proto.NewTemplateRequest,
) (*proto.NewTemplateResponse, error) {
	if in.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}
	if in.Description == "" {
		return nil, status.Error(codes.InvalidArgument, "description is required")
	}
	if in.attribute_id == [] {
		return nil, status.Error(codes.InvalidArgument, "attribute id is required")
	}

	err := s.template.NewTemplate(ctx, in.GetName(), in.GetDescription(), in.GetAttribute_id())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Errorf("failed to create new template: %w", err).Error())
	}

	return &proto.NewTemplateResponse{}, nil
}

func (s *ServerAPI) AlterTemplate(
	ctx context.Context,
	in *proto.AlterTemplateRequest,
) (*proto.AlterTemplateResponse, error) {
	if in.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}
	if in.Description == "" {
		return nil, status.Error(codes.InvalidArgument, "description is required")
	}
	if in.TemplateId == [] {
		return nil, status.Error(codes.InvalidArgument, "template id is required")
	}

	err := s.template.AlterTemplate(ctx, in.GetAttribute_id(), in.GetName(), in.GetDescription())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Errorf("failed to alter shop: %w", err).Error())
	}

	return &proto.AlterTemplateResponse{}, nil
}

func (s *ServerAPI) DeleteTemplate(
	ctx context.Context,
	in *proto.DeleteTemplateRequest,
) (*proto.DeleteTemplateResponse, error) {
	if in.Template_id == 0 {
		return nil, status.Error(codes.InvalidArgument, "template id is required")
	}

	err := s.template.DeleteTemplate(ctx, in.GetTemplate_id())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Errorf("failed to delete template: %w", err).Error())
	}

	return &proto.DeleteTemplateResponse{}, nil
}

func (s *ServerAPI) ListTemplates(
	ctx context.Context,
	in *proto.ListTemplatesRequest,
) (*proto.ListTemplatesResponse, error) {
	templateInfo, err := s.template.ListTemplates(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Errorf("failed to delete template: %w", err).Error())
	}

	return &proto.ListTemplatesResponse{Info: templateInfo}, nil
}
