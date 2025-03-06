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
	Template                          Template
}

// Тот самый интерфейс, котрый мы передавали в grpcApp
type Template interface {
	NewTemplate(
		ctx context.Context,
		name string,
		description string,
		branch_id int32,
		attributes []*proto.AttributeInfo,
	) error
	DeleteTemplate(
		ctx context.Context,
		template_id int32,
	) error
	ListTemplates(
		ctx context.Context,
		branch_id int32,
	) (
		[]*proto.TemplateInfo,
		error,
	)
}

func Register(gRPCServer *grpc.Server, template Template) {
	proto.RegisterTemplateServer(gRPCServer, &ServerAPI{Template: template})
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
	if in.BranchId == 0 {
		return nil, status.Error(codes.InvalidArgument, "branch id is required")
	}
	if len(in.Attributes) == 0 {
		return nil, status.Error(codes.InvalidArgument, "attribute structures is required")
	}
	if len(in.Attributes) > 10 {
		return nil, status.Error(codes.InvalidArgument, "too many attributes")
	}

	err := s.Template.NewTemplate(
		ctx,
		in.GetName(),
		in.GetDescription(),
		in.GetBranchId(),
		in.GetAttributes())

	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Errorf("failed to create new template: %w", err).Error())
	}

	return &proto.NewTemplateResponse{}, nil
}

func (s *ServerAPI) DeleteTemplate(
	ctx context.Context,
	in *proto.DeleteTemplateRequest,
) (*proto.DeleteTemplateResponse, error) {
	if in.TemplateId == 0 {
		return nil, status.Error(codes.InvalidArgument, "template id is required")
	}

	err := s.Template.DeleteTemplate(ctx, in.GetTemplateId())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Errorf("failed to delete template: %w", err).Error())
	}

	return &proto.DeleteTemplateResponse{}, nil
}

func (s *ServerAPI) ListTemplates(
	ctx context.Context,
	in *proto.ListTemplatesRequest,
) (*proto.ListTemplatesResponse, error) {
	if in.BranchId == 0 {
		return nil, status.Error(codes.InvalidArgument, "branch id is required")
	}
	templateInfo, err := s.Template.ListTemplates(ctx, in.GetBranchId())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Errorf("failed to list templates: %w", err).Error())
	}

	return &proto.ListTemplatesResponse{Info: templateInfo}, nil
}
