package template_app

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	proto "pim-sys/gen/go/template"
	auth_interceptor "pim-sys/internal/auth-interceptor"
	grpcapp "pim-sys/internal/grpc"
	template_service "pim-sys/internal/template/service"
	storage "pim-sys/internal/template/storage"

	"google.golang.org/grpc"
)

type App struct {
	GRPCServer *grpcapp.App
}

type Storage interface {
	CreateTemplate(
		ctx context.Context,
		branch_id int32,
		name string,
		description string,
		attributes []*proto.AttributeInfo,
	) error
	DeleteTemplate(
		ctx context.Context,
		templateId int32,
	) error
	ListTemplates(
		ctx context.Context,
		branch_id int32,
	) (
		[]*proto.TemplateInfo,
		error,
	)
	GetUserListBranches(
		ctx context.Context,
		user_id int32,
	) (
		[]int32,
		error,
	)
	GetBranchIdFromTemplateId(
		ctx context.Context,
		template_id int32,
	) (
		int32,
		error,
	)
}

type Template struct {
	TemplateStorage Storage
}

func (template *Template) NewTemplate(
	ctx context.Context,
	name string,
	description string,
	branch_id int32,
	attributes []*proto.AttributeInfo,
) error {
	return template.TemplateStorage.CreateTemplate(ctx, branch_id, name, description, attributes)
}

func (template *Template) DeleteTemplate(
	ctx context.Context,
	templateId int32,
) error {
	branchId, err := template.TemplateStorage.GetBranchIdFromTemplateId(ctx, templateId)
	if err != nil {
		return fmt.Errorf("%s: %v", "getting branch_id from template_id", err)
	}
	err = template.userMustHaveAccess(ctx, branchId)
	if err != nil {
		return fmt.Errorf("%s: %v", "checking user permissions", err)
	}

	return template.TemplateStorage.DeleteTemplate(ctx, templateId)
}

func (template *Template) ListTemplates(
	ctx context.Context,
	branch_id int32,
) (
	[]*proto.TemplateInfo,
	error,
) {

	err := template.userMustHaveAccess(ctx, branch_id)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", "checking user permissions", err)
	}

	return template.TemplateStorage.ListTemplates(ctx, int32(branch_id))
}

func New(
	log *slog.Logger,
	grpcPort int,
	connectionString string,
	tokenTTL time.Duration,
) *App {

	templateStorage, err := storage.New(connectionString)
	if err != nil {
		panic(err)
	}

	registerTemplate := func(gRPCServer *grpc.Server) {
		template_service.Register(
			gRPCServer,
			&Template{
				TemplateStorage: templateStorage,
			},
		)
	}

	grpcApp := grpcapp.New(log, registerTemplate, grpcPort, auth_interceptor.AuthInterceptor())

	return &App{
		GRPCServer: grpcApp,
	}
}

func (template *Template) userMustHaveAccess(ctx context.Context, branch_id int32) error {
	user_id, err := auth_interceptor.GetFromContext(ctx, "user_id")
	if err != nil {
		return fmt.Errorf("%s", "template operations: can't take current user id")
	}

	var user_idInt int32
	result, err := strconv.Atoi(user_id)
	user_idInt = int32(result)
	if err != nil {
		return fmt.Errorf("%s", "template operations: can't take cast user_id to int32")
	}
	availibleBranches, err := template.TemplateStorage.GetUserListBranches(ctx, user_idInt)
	if err != nil {
		return fmt.Errorf("%s%w", "template operations: can't take user's availible branches: ", err)
	}

	for _, category_id := range availibleBranches {
		if category_id == branch_id {
			return nil
		}
	}

	return fmt.Errorf("%s", "user access to branch denied or branch does not exist")
}
